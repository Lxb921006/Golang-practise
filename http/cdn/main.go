package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

var (
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

	urls    = flag.String("urls", "htpps://项目cdn域名.com/game/", "需要刷新的url")
	cdnType = flag.String("type", "dir", "刷新目录: dir, 刷新文件: file")
)

type ReflushCdn struct {
	uid       string
	accessKey string
	url       string
}

func (d *ReflushCdn) CreateNonce() string {
	timeNonce := strconv.Itoa(int(time.Now().Unix()))
	nonceByte := md5.Sum([]byte(timeNonce))
	noceSplice := nonceByte[0:]
	return string(noceSplice)[:11]
}

func (d *ReflushCdn) CreateSign() string {
	nonce := d.CreateNonce()
	sortString := []string{nonce, d.accessKey, strconv.Itoa(int(time.Now().Unix()))}
	sort.Strings(sortString)
	sortSign := strings.Join(sortString, "")

	h := sha1.New()
	h.Write([]byte(sortSign))
	return hex.EncodeToString(h.Sum(nil))
}

func (d *ReflushCdn) Reflush(params ...string) (resp []byte, err error) {
	var data = make(map[string]interface{})
	var headers = make(map[string]interface{})

	nonce := d.CreateNonce()
	sign := d.CreateSign()

	v := url.Values{}
	v.Set("sign", sign)
	v.Add("uid", d.uid)
	v.Add("nonce", nonce)
	v.Add("timestamp", strconv.Itoa(int(time.Now().Unix())))

	data["taskType"] = *cdnType
	data["urls"] = *urls

	url := d.url + "?" + v.Encode()

	nh := newHttp.NewHttpRe(url, data, headers, 4)
	resp, err = nh.POST(client)
	if err != nil {
		return
	}

	return
}

func NewReflushCdn() *ReflushCdn {
	return &ReflushCdn{
		uid:       "huawen",
		accessKey: "p802AULZXYcsl2vku5Xrnl3pKWieKzGdrPAaCfGJOQeKNK7hwzUgo6SKoIDQVpzA",
		url:       "https://openapi.wangjuyunlian.com/api/v1/content/purge",
	}
}

func main() {

	flag.Parse()

	if flag.NFlag() != 2 {
		log.Print(flag.ErrHelp.Error())
		return
	}

	nf := NewReflushCdn()
	resp, err := nf.Reflush()
	if err != nil {
		log.Println("刷新失败, esg >>>", err)
		return
	}

	log.Println("msg >>>", string(resp))
}
