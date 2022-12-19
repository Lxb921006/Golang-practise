package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

func main() {
	var sign string
	var nonce string
	var uid string
	var accessKey string
	var params = make(map[string]interface{})
	var headers = make(map[string]interface{})

	params["domain"] = "us-cdn-static.burstedgold.com"
	params["start"] = "2022-12-15"
	params["end"] = "2022-12-17"

	timeNonce := strconv.Itoa(int(time.Now().Unix()))
	nonceByte := md5.Sum([]byte(timeNonce))
	noceSplice := nonceByte[0:]
	nonce = string(noceSplice)[:11]

	uid = "maiyou"
	accessKey = "XcaLL4fMkkSsnIcoyhq6aSFC8QXKkKpo0rYI3TvaGutjF70blSRZrXpzw0PSrGu4"

	sortString := []string{nonce, accessKey, strconv.Itoa(int(time.Now().Unix()))}
	sort.Strings(sortString)
	sortSign := strings.Join(sortString, "")

	h := sha1.New()
	h.Write([]byte(sortSign))
	sign = hex.EncodeToString(h.Sum(nil))

	v := url.Values{}
	v.Set("sign", sign)
	v.Add("uid", uid)
	v.Add("nonce", nonce)
	v.Add("timestamp", strconv.Itoa(int(time.Now().Unix())))
	splice := v.Encode()

	url := "https://openapi.wangjuyunlian.com/api/v1/log/list?" + splice

	nh := newHttp.NewHttpRe(url, params, headers, 4)
	data, err := nh.GET()
	if err != nil {
		log.Print(err)
		return
	}

	log.Print("data = ", string(data))
}
