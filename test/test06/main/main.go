package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	m1 := map[string]interface{}{}
	httpUrl := "http://101.35.143.86/login/"
	data := url.Values{
		"username": {"lxb"},
		"password": {"7109667@Lxb"},
	}

	resp, err := http.PostForm(httpUrl, data)
	resp.Header.Set("Content-Type", "application/json")

	defer resp.Body.Close()
	if err != nil {
		fmt.Println("request err = ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body err = ", err)
		return
	}

	err = json.Unmarshal(body, &m1)
	if err != nil {
		fmt.Println("parse err = ", err)
		return
	}

	fmt.Printf("res = %v\n", m1["msg"])
}
