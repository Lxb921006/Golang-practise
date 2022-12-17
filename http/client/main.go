package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var data map[string]interface{}
	var params = make(map[string]interface{})
	v := url.Values{}

	request_url := "http://api.btstu.cn/qqol/api.php?qq=10001"
	// request_url := "https://api.xwteam.cn/api/qq/music"

	params["qq"] = "120332269"
	// params["password"] = "12332"

	b, _ := json.Marshal(&params)

	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}

	// hr, err := http.NewRequest("GET", request_url, bytes.NewReader(b))

	v.Set("qq", "120332269")
	// v.Set("password", "12332")
	// p := v.Encode()

	hr, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Print("req error=>", err)
		return
	}

	hr.Header.Add("content-type", "application/x-www-form-urlencoded")

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
