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
	"time"
)

func main() {
	httpServer()
}

type Resp struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Date   string `json:"date"`
}

func (r *Resp) M(msg string, code int) (b []byte) {
	r.Msg = msg
	r.Status = code
	r.Date = time.Now().Format("2006-01-02 15:04:05")
	b, _ = json.Marshal(r)

	return
}

func httpServer() {
	log.Println("http server :8092 listening...")

	mux := http.NewServeMux()
	mux.HandleFunc("/download", download)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/content", sendFileContent)
	mux.HandleFunc("/aws-cdn-refresh", awsCdnRefresh)

	listen := &http.Server{
		Addr:              ":8092",
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(10) * time.Second,
	}

	if err := listen.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
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

	// 获取额外的参数
	value := request.Form.Get("user")
	fmt.Println(value)

	saveDir := filepath.Join("C:\\Users\\Administrator\\Desktop\\test", header.Filename)
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
	fp := filepath.Join("C:\\Users\\Administrator\\Desktop", file)
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
	var ctx = context.Background()
	if request.Method != "GET" {
		b := resp.M("请求方法错误", 10003)
		writer.Write(b)
		return
	}

	var f = request.URL.Query()
	var path = f.Get("path")
	if path == "" {
		b := resp.M("刷新目录不能为空", 10002)
		writer.Write(b)
		return
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(10))
	defer cancel()

	out, err := exec.CommandContext(ctx, "sh", "/root/aws_cdn_refresh.sh", path).Output()
	fmt.Printf("---------%s----------\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(string(out))
	if err != nil {
		b := resp.M(string(out), 10001)
		writer.Write(b)
		return
	}

	b := resp.M(fmt.Sprintf("%s刷新成功, 刷新生效需等1分钟左右", path), 10000)

	writer.Write(b)

}
