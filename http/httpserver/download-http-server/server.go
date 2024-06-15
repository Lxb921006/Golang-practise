package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	httpServer()
}

type Resp struct {
	Msg    string                   `json:"msg"`
	Status int                      `json:"status"`
	Date   string                   `json:"date"`
	Detail map[string]interface{}   `json:"detail,omitempty"`
	Data   []map[string]interface{} `json:"data,omitempty"`
	Br     []byte                   `json:"br,omitempty"`
}

func (r *Resp) M(msg string, code int) (b []byte) {
	r.Msg = msg
	r.Status = code
	r.Date = time.Now().Add(time.Hour * time.Duration(11)).Format("2006-01-02 15:04:05")
	b, _ = json.Marshal(r)

	return
}

func (r *Resp) K(resp *Resp) (b []byte) {
	r.Msg = resp.Msg
	r.Status = resp.Status
	r.Date = time.Now().Add(time.Hour * time.Duration(11)).Format("2006-01-02 15:04:05")
	r.Detail = resp.Detail
	b, _ = json.Marshal(r)

	return
}

func (r *Resp) R(writer http.ResponseWriter) error {
	_, err := writer.Write(r.Br)
	if err != nil {
		return err
	}
	return nil
}

func httpServer() {
	log.Println("http server :8092 listening...")

	mux := http.NewServeMux()
	mux.HandleFunc("/download", download)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/content", sendFileContent)
	mux.HandleFunc("/aws-cdn-refresh", awsCdnRefresh)
	mux.HandleFunc("/wx-data", wxGetData)
	mux.HandleFunc("/change-time", changeTime)
	listen := &http.Server{
		Addr:              ":8092",
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(10) * time.Second,
	}

	if err := listen.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
}

func changeTime(response http.ResponseWriter, request *http.Request) {
	var resp Resp
	if request.Method != "GET" {
		b := resp.M("请求方法错误", 10003)
		if _, err := response.Write(b); err != nil {
			return
		}

		return
	}

	p := request.URL.Query()
	date := p.Get("date")
	sign := p.Get("sign")

	if sign != "change" {
		b := resp.M("无效请求", 10004)
		if _, err := response.Write(b); err != nil {
			return
		}

		return
	}

	log.Printf("date = %s, sign = %s\n", date, sign)

	scriptPath := "/web/wwwroot/777brs.com/web/change_time.sh"
	output, err := exec.Command("/bin/bash", scriptPath, date).Output()
	if err != nil {
		b := resp.M(fmt.Sprintf("执行%s失败，失败信息: %s", scriptPath, err), 10005)
		if _, err := response.Write(b); err != nil {
			return
		}

		return
	}

	b := resp.M(fmt.Sprintf("执行%s成功，成功信息: %s", scriptPath, string(output)), 10000)
	if _, err := response.Write(b); err != nil {
		return
	}

}

func wxGetData(resp http.ResponseWriter, req *http.Request) {
	var data = []map[string]interface{}{
		{
			"id":   1,
			"name": "lxb",
			"age":  31,
		},
		{
			"id":   2,
			"name": "lqm",
			"age":  18,
		},
		{
			"id":   3,
			"name": "lyy",
			"age":  17,
		},
	}

	var r Resp
	if req.Method != "GET" {
		b := r.M("请求方法错误", 10003)
		resp.Write(b)
		return
	}

	log.Println(req.RequestURI)

	f := req.URL.Query()
	var num = f.Get("num")
	li, _ := strconv.Atoi(num)
	if li < len(data) {
		data = data[:li]
	}

	r = Resp{
		Msg:    "ok",
		Status: 10000,
		Data:   data,
	}
	b := r.K(&r)

	resp.Write(b)

}

func upload(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	if request.Method != "POST" {
		b := resp.M("请求方法错误", 10003)
		writer.Write(b)
		return
	}

	form := request.ParseMultipartForm(32 << 20)
	if form != nil {
		b := resp.M(form.Error(), 10001)
		writer.Write(b)
		return
	}

	file, header, _ := request.FormFile("file")

	fmt.Println("header >>> ", header)
	// 获取额外的参数
	value := request.Form.Get("user")
	fileName := request.Form.Get("fileName")
	fmt.Println(value)

	saveDir := filepath.Join("/nas/th-db-bak", fileName)
	fc, _ := os.Create(saveDir)

	defer fc.Close()

	_, err := io.Copy(fc, file)
	if err != nil {
		b := resp.M(err.Error(), 10001)
		writer.Write(b)
		return
	}

	b := resp.M("上传成功", 10000)
	writer.Write(b)

}

func download(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	if request.Method != "GET" {
		b := resp.M("请求方法错误", 10003)
		writer.Write(b)
		return
	}

	f := request.URL.Query()
	if f.Get("file") == "" {
		b := resp.M("file字段不能为空", 10001)
		writer.Write(b)
		return
	}

	if err := sendFileHandle(f.Get("file"), writer); err != nil {
		b := resp.M(err.Error(), 10002)
		writer.Write(b)
		return
	}

	b := resp.M("ok", 10000)
	writer.Write(b)

}

func sendFileHandle(file string, w http.ResponseWriter) (err error) {
	fp := filepath.Join("D:\\soft", file)
	f, err := os.Open(fp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file))
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Header().Set("Content-Length", "123456789") // 设置文件大

	_, err = io.Copy(w, f)
	if err != nil {
		return
	}

	return
}

func sendFileContent(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	if request.Method != "GET" {
		b := resp.M("请求方法错误", 10003)
		writer.Write(b)
		return
	}

	f := request.URL.Query()
	file := filepath.Join("C:\\Users\\Administrator\\Desktop", f.Get("file"))
	_, err := os.Stat(file)
	if err != nil {
		b := resp.M(err.Error(), 10002)
		writer.Write(b)
		return
	}

	http.ServeFile(writer, request, file)

}

func awsCdnRefresh(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	var awsResp map[string]interface{}
	if request.Method != "GET" {
		b := resp.M("请求方法错误", 10003)
		_, _ = writer.Write(b)

		return
	}

	f := request.URL.Query()
	path := f.Get("path")
	item := f.Get("item")
	if path == "" || item == "" {
		b := resp.M("刷新目录或项目名不能为空", 10002)
		writer.Write(b)
		return
	}

	var cancel context.CancelFunc
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(10))
	defer cancel()

	out, err := exec.CommandContext(ctx, "sh", "/root/shellscript/aws_cdn_refresh.sh", item, path).Output()
	if err != nil {
		b := resp.M(string(out), 10001)
		writer.Write(b)
		return
	}

	err = json.Unmarshal(out, &awsResp)
	if err != nil {
		b := resp.M(err.Error(), 10004)
		writer.Write(b)
		return
	}

	respKRM := &Resp{
		Msg:    fmt.Sprintf("%s刷新成功, 刷新生效需等20-50s左右", path),
		Status: 10000,
		Detail: awsResp,
	}

	b := resp.K(respKRM)
	respKRM.Br = b

	respKRM.R(writer)
}
