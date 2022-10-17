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

//application/json: django为例直接request.post是获取不了数据, 需要request.Body去获取
type ParseParams struct {
}

func (p *ParseParams) FormatParams(data string) (ns []byte) {
	requestStruct := map[string]interface{}{}
	b, _ := json.Marshal(&data)
	json.Unmarshal(b, &requestStruct)
	ns, _ = json.Marshal(&requestStruct)
	return
}

func (p *ParseParams) PostRequest(httpUrl, header, params string) (res map[string]interface{}, err error) {
	var request *http.Request
	switch header {
	case "application/json":
		b := p.FormatParams(params)
		request, err = http.NewRequest("POST", httpUrl, bytes.NewReader(b))
		request.Header.Set("Content-Type", header)
		if err != nil {
			err = fmt.Errorf("create request err = %v", err)
			return
		}
	case "application/x-www-form-urlencoded", "multipart/form-data":
		request, err = http.NewRequest("POST", httpUrl, strings.NewReader(params))
		request.Header.Set("Content-Type", header)
		if err != nil {
			err = fmt.Errorf("create request err = %v", err)
			return
		}
	default:
		err = errors.New("invalid header")
		return
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("request err = %v", err)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read err = %v", err)
		return
	}

	json.Unmarshal(data, &res)
	return
}

func (p *ParseParams) GetRequest(httpUrl, header, params string) (res map[string]interface{}, err error) {
	return
}

func main() {
	httpUrl := "http://data.burstedgold.com:8992/open/get-sql-for-user-search?"
	token := "Ld86A1BLHffMyyJglHmumYJnSjJ8hfmZORoCtcuDHoLl7eOoVrVRhX2xiNN3sQRp"
	projectId := 123
	httpUrl = fmt.Sprintf("%stoken=%s&projectId=%d", httpUrl, token, projectId)
	params := `{
        "filter": {
            "filterType": "COMPOUND",
            "relation": "and",
            "filts": [
                {
                    "filterType": "SIMPLE",
                    "tableType": "user",
                    "columnName": "part_event",
                    "comparator": "equal",
                    "ftv": [
                        "login"
                    ]
                },
                {
                    "filterType": "SIMPLE",
                    "tableType": "user",
                    "columnName": "part_date",
                    "comparator": "greater than",
                    "ftv": [
                        "2021-06-20"
                    ]
                }
            ]
        },
        "selectAllColumns": "false",
        "selectColumns": [
            "#count()"
        ]
    }`
	postApi := &ParseParams{}
	//封装了NewRequest post方法, arg1:url,arg2:header,arg3:params
	resp, err := postApi.PostRequest(httpUrl, "application/json", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

}
