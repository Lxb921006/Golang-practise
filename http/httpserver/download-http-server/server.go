package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	httpServer()
}

type Resp struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

func (r *Resp) M(msg string, code int) (b []byte) {
	r.Msg = msg
	r.Status = code
	b, _ = json.Marshal(r)
	return
}

func httpServer() {
	log.Println("http server 8092 listening...")

	mux := http.NewServeMux()
	mux.Handle("/download", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var resp Resp
		if request.Method == "GET" {
			f := request.URL.Query()
			if f.Get("file") == "" {
				b := resp.M("file字段不能为空", 10001)
				writer.Write(b)
			} else {
				if err := sendFileHandle(f.Get("file"), writer); err != nil {
					b := resp.M(err.Error(), 10001)
					writer.Write(b)
				} else {
					b := resp.M("ok", 10000)
					writer.Write(b)
				}
			}
		}

	}))

	listen := &http.Server{
		Addr:              ":8092",
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(10) * time.Second,
	}

	if err := listen.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

}

func sendFileHandle(file string, w http.ResponseWriter) (err error) {
	//var nb = make([]byte, 1024)
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
