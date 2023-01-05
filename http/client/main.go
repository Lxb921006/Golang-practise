package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var data map[string]interface{}
	var params = make(map[string]interface{})
	// v := url.Values{}

	request_url := "http://47.241.38.210:9294/cron/run?user=lxb&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Imx4YiIsImV4cCI6MTY3MzE0OTQ5MH0.VqeY_6H08HBsaALy-lrp_GkFXLDSxBfS1zOC5-V6tUk"
	// request_url := "https://api.xwteam.cn/api/qq/music"

	params["cron_id"] = 1162
	params["crons"] = "/usr/local/php/bin/php -f /web/wwwroot/shell/thailand_burstedgold_cron/index_cli.php Shell/Gameonlinedata/index"

	b, _ := json.Marshal(&params)

	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}

	hr, err := http.NewRequest("POST", request_url, bytes.NewReader(b))

	// v.Set("cron_id", "1162")
	// v.Set("password", "12332")
	// p := v.Encode()

	// hr, err := http.NewRequest("POST", request_url, strings.NewReader(p))
	if err != nil {
		log.Print("req error=>", err)
		return
	}

	// hr.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(hr)
	if err != nil {
		log.Print("resp error=>", err)
		return
	}

	defer resp.Body.Close()

	b, _ = ioutil.ReadAll(resp.Body)
	log.Print(string(b))

	if resp.StatusCode != 200 {
		log.Print("resp error=>", string(b))
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Print("serialize error=>", err)

	} else {
		log.Print(string(b))
	}

	// if resp.StatusCode != 200 {
	// 	log.Print(data["message"])
	// }

	log.Print(data)
}
