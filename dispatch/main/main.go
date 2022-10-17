package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//反射入门案列
	p := Person{
		Name: "lxb",
		Age:  30,
	}

	b, _ := json.Marshal(&p)
	fmt.Println(string(b))
}
