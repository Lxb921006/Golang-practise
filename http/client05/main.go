package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

var (
	wg sync.WaitGroup
)

type DownloadLog struct {
	uid            string
	accessKey      string
	url            string
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
	resp, err = nh.GET()
	if err != nil {
		return
	}

	return
}

func (d *DownloadLog) DownloadToLocal(params ...string) (err error) {
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

	const MaxWorkers = 20
	wg.Add(20)
	// downloadPath1 := "C:/Users/Administrator/Desktop/log"
	downloadPath1 := "/Users/liaoxuanbiao/Downloads/log"

	// for i := 0; i < MaxWorkers; i++ {
	// 	go d.UnGz()
	// }

	// 20个下载goroutine
	for i := 0; i < MaxWorkers; i++ {
		go d.WriteToFile(downloadPath1)
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
	//到这里可以确保不再往管道发送数据，可以关闭，遵循关闭管道原则
	close(d.downloadWork)
	return
}

func (d *DownloadLog) WriteToFile(path string) {
	// defer func() { d.toStop <- "stop"; log.Print("stoppppppppppp") }() //这里是确保所有文件下载并解压完成后发送信号给toStop停止接受数据
	defer wg.Done()
	for v := range d.downloadWork {
		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		url := strings.Split(v, "+")[0]
		date := strings.Split(v, "+")[1]
		urlGz := strings.Split(url, "?")[0]
		filename := filepath.Join(filepath.Join(path, date), filepath.Base(urlGz))
		log.Print("goroutine number = ", runtime.NumGoroutine(), ",  url = ", filename)

		var data = make(map[string]interface{})
		var headers = make(map[string]interface{})
		nh := newHttp.NewHttpRe(url, data, headers, 4)
		nh.GET()

	}
}

func (d *DownloadLog) UnGz() {
	defer func() { wg.Done(); log.Print("recv stopppppp") }()

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
			time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
			log.Print("recv == ", file)
			// unGz := extract.NewUngz(file)
			// err := unGz.UngzFile()
			// if err != nil {
			// 	d.ungzFailed <- file
			// }
		}
	}
}

func NewDownloadLog(url string) *DownloadLog {
	return &DownloadLog{
		uid:            "maiyou",
		accessKey:      "XcaLL4fMkkSsnIcoyhq6aSFC8QXKkKpo0rYI3TvaGutjF70blSRZrXpzw0PSrGu4",
		url:            url,
		downloadWork:   make(chan string),
		downloadFailed: make(chan string),
		ungzWork:       make(chan string),
		ungzFailed:     make(chan string),
		toStop:         make(chan string, 1),
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	start := time.Now()
	url := "https://openapi.wangjuyunlian.com/api/v1/log/list?"
	nd := NewDownloadLog(url)
	err := nd.DownloadToLocal("us-cdn-static.burstedgold.com", "2022-12-15", "2022-12-15")
	if err != nil {
		log.Print(err)
		return
	}
	wg.Wait()
	log.Print("cost time = ", time.Since(start))

	close(nd.downloadFailed)
	close(nd.ungzFailed)
}
