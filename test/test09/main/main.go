package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestPkg struct {
}

//这个是给对应请求头类型参数进行格式化
func (r *RequestPkg) FormatParams(data string) (ns []byte) {
	var requestStruct map[string]interface{}
	b, _ := json.Marshal(&data)
	json.Unmarshal(b, &requestStruct)
	ns, _ = json.Marshal(&requestStruct)
	return
}

func (r *RequestPkg) PostRequest(httpUrl, header, params string) (res map[string]interface{}, err error) {
	var resp *http.Response

	switch header {
	case "application/json":
		b := r.FormatParams(params)
		resp, err = http.Post(httpUrl, header, bytes.NewReader(b))
		if err != nil {
			return res, err
		}
	case "application/x-www-form-urlencoded", "multipart/form-data":
		resp, err = http.Post(httpUrl, header, strings.NewReader(params))
		if err != nil {
			return res, err
		}
	default:
		err = errors.New("invalid header")
		return
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("request err = ", err)
		return
	}

	return
}

func (r *RequestPkg) GetRequest(httpUrl string) (res map[string]interface{}, err error) {
	resp, err := http.Get(httpUrl)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("request err = ", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read err = %v", err)
		return
	}

	json.Unmarshal(data, &res)

	return
}

func main() {
	r := &RequestPkg{}
	httpUrl := "http://127.0.0.1:9293/login"
	data := "user=lxb&password=123322"
	header := "application/x-www-form-urlencoded"

	resp, err := r.PostRequest(httpUrl, header, data)
	if err != nil {
		fmt.Println("request err = ", err)
		return
	}
	fmt.Println("resp = ", resp)
}
