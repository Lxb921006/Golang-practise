package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

type DownloadLog struct {
	uid       string
	accessKey string
	url       string
	work      chan string
	failed    chan string
	Maxworker int
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
	downloadPath1 := "C:/Users/Administrator/Desktop/log"

	go func() {
		for v := range <-d.failed {
			log.Print("download failed = ", v)
		}
	}()

	for range [MaxWorkers]struct{}{} {
		go d.WriteToFile(downloadPath1)
	}

	for i := 0; i < len(data["data"].([]interface{})); i++ {
		date := data["data"].([]interface{})[i].(map[string]interface{})["date"].(string)
		downloadPath2 := filepath.Join(downloadPath1, date)
		_, err = os.Stat(downloadPath2)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(downloadPath2, 0777)
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
					d.work <- data
					break
				}
			}
		}
	}

	return
}

func (d *DownloadLog) WriteToFile(path string) {
	for v := range d.work {
		url := strings.Split(v, "+")[0]
		date := strings.Split(v, "+")[1]
		urlGz := strings.Split(url, "?")[0]
		filename := filepath.Join(filepath.Join(path, date), filepath.Base(urlGz))

		var data = make(map[string]interface{})
		var headers = make(map[string]interface{})

		nh := newHttp.NewHttpRe(url, data, headers, 4)
		resp, err := nh.GET()
		if err == nil {
			fn, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0777)
			if err == nil {
				fn.WriteString(string(resp))
			} else {
				d.failed <- url
			}
		} else {
			d.failed <- url
		}
	}
}

func NewDownloadLog(url string) *DownloadLog {
	return &DownloadLog{
		uid:       "maiyou",
		accessKey: "XcaLL4fMkkSsnIcoyhq6aSFC8QXKkKpo0rYI3TvaGutjF70blSRZrXpzw0PSrGu4",
		url:       url,
		work:      make(chan string),
	}
}

func main() {
	url := "https://openapi.wangjuyunlian.com/api/v1/log/list?"
	err := NewDownloadLog(url).DownloadToLocal("us-cdn-static.burstedgold.com", "2022-12-15", "2022-12-17")
	if err != nil {
		log.Print(err)
		return
	}

}
