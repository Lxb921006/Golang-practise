package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//x-www-form-urlencoded:会把表单转成键值对如：username=lxb&password=123

func main() {
	respStruct := map[string]interface{}{}
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
	requestStruct := map[string]interface{}{}
	b, _ := json.Marshal(&params)
	json.Unmarshal(b, &requestStruct)
	ns, _ := json.Marshal(&requestStruct)
	//参数1:url,2:请求头类型,3:切片
	resp, err := http.Post(httpUrl, "application/json", bytes.NewReader(ns))
	fmt.Println(resp.StatusCode)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("request err = ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err = ", err)
		return
	}

	json.Unmarshal(body, &respStruct)

	fmt.Println("body = ", respStruct)

}
