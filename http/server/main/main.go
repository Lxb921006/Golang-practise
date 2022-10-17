package main

import (
	"fmt"
	"html"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

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
