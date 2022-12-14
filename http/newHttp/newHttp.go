package newHttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpRe struct {
	Url     string                 `json:"url"`
	Params  map[string]interface{} `json:"params"`
	Headers map[string]interface{} `json:"headers"`
	Timeout int                    `json:"timeout"`
	client  *http.Client
	hr      *http.Request
	resp    *http.Response
}

func (nh *HttpRe) POST(client *http.Client) (data []byte, err error) {
	fd, err := nh.FormatParams()
	if err != nil {
		return
	}

	nh.client = client

	data, err = nh.NewRequest("POST", fd.(io.Reader))
	if err != nil {
		return
	}

	defer nh.resp.Body.Close()

	return
}

func (nh *HttpRe) GET(client *http.Client) (data []byte, err error) {
	fd, err := nh.FormatParams()
	if err != nil {
		return
	}

	nh.client = client

	data, err = nh.NewRequest("GET", fd.(io.Reader))

	if err != nil {
		return
	}

	defer nh.resp.Body.Close()

	return
}

func (nh *HttpRe) FormatParams() (data interface{}, err error) {
	switch nh.Headers["content-type"] {
	case "application/x-www-form-urlencoded":
		vv := url.Values{}
		for k, v := range nh.Params {
			vv.Set(k, v.(string))
		}

		data = strings.NewReader(vv.Encode())
	case "application/json":
		b, errs := json.Marshal(&nh.Params)
		if errs != nil {
			err = fmt.Errorf("序列化参数错误, %v", errs)
			return
		}

		data = bytes.NewReader(b)
	default:
		nh.Headers["content-type"] = "application/json"
		b, errs := json.Marshal(&nh.Params)
		if errs != nil {
			err = fmt.Errorf("序列化参数错误, %v", errs)
			return
		}

		data = bytes.NewReader(b)
	}

	return

}

func (nh *HttpRe) NewRequest(method string, params io.Reader) (data []byte, err error) {
	switch method {
	case "POST":
		nh.hr, err = http.NewRequest("POST", nh.Url, params)
		if err != nil {
			err = fmt.Errorf("创建POST请求失败, %v", err)
			return
		}
	case "GET":
		nh.hr, err = http.NewRequest("GET", nh.Url, params)
		if err != nil {
			err = fmt.Errorf("创建GET请求失败, %v", err)
			return
		}
	default:
		err = fmt.Errorf("方法不存在, %v", err)
		return
	}

	//请求头
	nh.hr.Header.Add("content-type", nh.Headers["content-type"].(string))

	//响应
	nh.resp, err = nh.client.Do(nh.hr)
	if err != nil {
		err = fmt.Errorf("请求失败, esg = %v", err)
		return
	}

	data, err = io.ReadAll(nh.resp.Body)
	if err != nil {
		err = fmt.Errorf("获取响应数据失败, esg = %v", err)
		return
	}

	//状态码
	if nh.resp.StatusCode != 200 {
		err = errors.New(string(data))
		return
	}

	return
}

func NewHttpRe(url string, params, headers map[string]interface{}, tt int) *HttpRe {
	return &HttpRe{
		Url:     url,
		Params:  params,
		Headers: headers,
		Timeout: 5,
	}
}
