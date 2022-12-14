package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Lxb921006/Golang-practise/extract"
	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

var (
	wg     sync.WaitGroup
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 3 * time.Second,
		},
		Timeout: time.Duration(10) * time.Second,
	}
)

type DownloadLog struct {
	uid            string
	accessKey      string
	url            string
	Path           string
	downloadWork   chan string
	ungzWork       chan string
	downloadFailed chan string
	ungzFailed     chan string
	toStop         chan string
	Maxworker      int
}

func (d *DownloadLog) CreateNonce() string {
	timeNonce := strconv.Itoa(int(time.Now().Unix()))
	nonceByte := md5.Sum([]byte(timeNonce))
	noceSplice := nonceByte[0:]
	return string(noceSplice)[:11]
}

func (d *DownloadLog) CreateSign() string {
	nonce := d.CreateNonce()
	sortString := []string{nonce, d.accessKey, strconv.Itoa(int(time.Now().Unix()))}
	sort.Strings(sortString)
	sortSign := strings.Join(sortString, "")

	h := sha1.New()
	h.Write([]byte(sortSign))
	return hex.EncodeToString(h.Sum(nil))
}

func (d *DownloadLog) RequestApi(params ...string) (resp []byte, err error) {
	var data = make(map[string]interface{})
	var headers = make(map[string]interface{})

	nonce := d.CreateNonce()
	sign := d.CreateSign()

	v := url.Values{}
	v.Set("sign", sign)
	v.Add("uid", d.uid)
	v.Add("nonce", nonce)
	v.Add("domain", params[0])
	v.Add("start", params[1])
	v.Add("end", params[2])
	v.Add("timestamp", strconv.Itoa(int(time.Now().Unix())))

	url := d.url + v.Encode()
	nh := newHttp.NewHttpRe(url, data, headers, 4)
	resp, err = nh.GET(client)
	if err != nil {
		return
	}

	return
}

func (d *DownloadLog) DownloadToLocal(params ...string) (err error) {
	const MaxWorkers = 20
	wg.Add(20)
	downloadPath1 := params[3]
	var data = make(map[string]interface{})

	resp, err := d.RequestApi(params[0], params[1], params[2])
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return
	}

	if data["code"].(string) != "000000" {
		err = errors.New(data["msg"].(string))
		return
	}

	//?????????????????????url
	go func() {
		for v := range <-d.downloadFailed {
			log.Print("download failed = ", v)
		}
	}()

	//???????????????????????????
	go func() {
		for v := range <-d.ungzFailed {
			log.Print("ungz failed = ", v)
		}
	}()

	// 20?????????goroutine
	for range [MaxWorkers]struct{}{} {
		go d.WriteToFile(downloadPath1)
	}

	// 20?????????goroutine
	for range [MaxWorkers]struct{}{} {
		go d.UnGz()
	}

	for i := 0; i < len(data["data"].([]interface{})); i++ {
		date := data["data"].([]interface{})[i].(map[string]interface{})["date"].(string)
		downloadPath2 := filepath.Join(downloadPath1, date)
		_, err = os.Stat(downloadPath2)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(downloadPath2, 0777)
				if err != nil {
					return
				}
			}
		}
		for k1, v1 := range data["data"].([]interface{})[i].(map[string]interface{}) {
			if k1 == "urls" {
				for _, v2 := range v1.([]interface{}) {
					url := v2.(string)
					data := url + "+" + date
					d.downloadWork <- data
				}
			}
		}
	}
	//??????????????????????????????????????????????????????????????????????????????????????????
	close(d.downloadWork)
	return
}

func (d *DownloadLog) WriteToFile(path string) {
	defer func() { d.toStop <- "stop" }() //??????????????????????????????????????????????????????????????????toStop??????????????????

	for v := range d.downloadWork {
		url := strings.Split(v, "+")[0]
		date := strings.Split(v, "+")[1]
		urlGz := strings.Split(url, "?")[0]
		filename := filepath.Join(filepath.Join(path, date), filepath.Base(urlGz))

		var data = make(map[string]interface{})
		var headers = make(map[string]interface{})
		nh := newHttp.NewHttpRe(url, data, headers, 4)
		resp, err := nh.GET(client)
		if err == nil {
			fn, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0777)
			if err == nil {
				//???????????????
				_, err = fn.WriteString(string(resp))
				if err == nil {
					//??????
					d.ungzWork <- filename
				} else {
					d.downloadFailed <- url
				}
			} else {
				d.downloadFailed <- url
			}
		} else {
			d.downloadFailed <- url
		}
	}
}

func (d *DownloadLog) UnGz() {
	defer wg.Done()

	for {
		select {
		case <-d.toStop:
			return
		default:
		}

		select {
		case <-d.toStop:
			return
		case file := <-d.ungzWork:
			unGz := extract.NewUngz(file)
			err := unGz.UngzFile()
			if err != nil {
				d.ungzFailed <- file
			}
		}
	}
}

func NewDownloadLog() *DownloadLog {
	return &DownloadLog{
		uid:            "maiyou",
		accessKey:      "XcaLL4fMkkSsnIcoyhq6aSFC8QXKkKpo0rYI3TvaGutjF70blSRZrXpzw0PSrGu4",
		url:            "https://openapi.wangjuyunlian.com/api/v1/log/list?",
		downloadWork:   make(chan string),
		downloadFailed: make(chan string),
		ungzWork:       make(chan string),
		ungzFailed:     make(chan string),
		toStop:         make(chan string, 1),
	}
}

func main() {

	domain := flag.String("domain", "", "??????")
	start1 := flag.String("start", "", "????????????")
	end := flag.String("end", "", "????????????")
	path := flag.String("path", "", "??????????????????")

	flag.Parse()

	if flag.NFlag() != 4 {
		log.Print(flag.ErrHelp)
		return
	}

	start := time.Now()
	nd := NewDownloadLog()
	err := nd.DownloadToLocal(
		*domain,
		*start1,
		*end,
		*path,
	)
	if err != nil {
		log.Print(err)
		return
	}
	wg.Wait()
	log.Print("cost time = ", time.Since(start))

	close(nd.downloadFailed)
	close(nd.ungzFailed)
}
