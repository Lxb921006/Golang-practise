package main

import "fmt"

type User struct {
	Name string
}

func main() {
	var m = make(map[string]interface{})

	// m["name"] = "lxb"

	v, ok := m["name"]
	if !ok {
		fmt.Println("null")
	} else {
		fmt.Println(v)
	}
}
