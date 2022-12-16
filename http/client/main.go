package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	var data map[string]interface{}
	var params = make(map[string]interface{})
	v := url.Values{}

	request_url := "http://127.0.0.1:9293/login"
	// request_url := "https://api.xwteam.cn/api/qq/music"

	params["user"] = "lxb"
	params["password"] = "12332"

	b, _ := json.Marshal(&params)
	strings.NewReader(string(b))

	client := &http.Client{
		Timeout: time.Duration(4) * time.Second,
	}
	// hr, err := http.NewRequest("POST", request_url, bytes.NewReader(b))

	v.Set("user", "lxb")
	v.Set("password", "12332")
	p := v.Encode()

	hr, err := http.NewRequest("POST", request_url, strings.NewReader(p))
	if err != nil {
		log.Print("req error=>", err)
		return
	}

	hr.Header.Add("content-type", "application/x-www-form-urlencoded")

	defer hr.Body.Close()

	resp, err := client.Do(hr)
	if err != nil {
		log.Print("resp error=>", err)
		return
	}

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
