package main

import "fmt"

type Public struct {
	Name string
	Age  int
}

type User struct {
	Public
	Skill string
}

func main() {
	u := User{
		Skill: "篮球",
		Public: Public{
			Name: "lxb",
			Age:  30,
		},
	}
	fmt.Println("u = ", u)
}
