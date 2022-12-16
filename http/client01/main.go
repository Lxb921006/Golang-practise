package main

import (
	"log"

	"github.com/Lxb921006/Golang-practise/http/newHttp"
)

func main() {
	var params = make(map[string]interface{})
	var headers = make(map[string]interface{})
	url := "http://127.0.0.1:9293/login"
	timeout := 5
	params["user"] = "lxb1"
	params["password"] = "123321"
	// headers["content-type"] = "application/x-www-form-urlencoded"

	nh := newHttp.NewHttpRe(url, params, headers, timeout)
	data, err := nh.POST()
	if err != nil {
		log.Print("请求失败, ", err)
		return
	}

	log.Print("data = ", string(data))
}
