package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 15)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))

}

func main() {
	fmt.Println("-------http服务端已开启-------")
	http.HandleFunc("/hello", SayHello)
	err := http.ListenAndServe(":10086", nil)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

}
