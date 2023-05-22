package main

import (
	"log"
	"net/http"
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
)

func main() {
	var params = make(map[string]interface{})
	var headers = make(map[string]interface{})
	url := "http://47.241.38.210:9294/cron/run?user=lxb&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Imx4YiIsImV4cCI6MTY3MzE0OTQ5MH0.VqeY_6H08HBsaALy-lrp_GkFXLDSxBfS1zOC5-V6tUk"
	url = "http://127.0.0.1:8000"
	timeout := 5
	params["cron_id"] = "1162"
	params["crons"] = "/usr/local/php/bin/php -f /web/wwwroot/shell/thailand_burstedgold_cron/index_cli.php Shell/Gameonlinedata/index"
	headers["content-type"] = "application/x-www-form-urlencoded"

	nh := newHttp.NewHttpRe(url, params, headers, timeout)
	//data, err := nh.POST(client)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}

	data, err := nh.GET(client)
	if err != nil {
		log.Print(err)
		return
	}

	log.Print("data = ", string(data))
}
