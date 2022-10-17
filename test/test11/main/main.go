package main

import "fmt"

type Humbeing struct {
	Name   string
	Gender string
}

type User struct {
	Humbeing
	Age int
}

func (u User) PrintAge() {
	fmt.Println("age = ", u.Age)
}

func main() {
	u := User{
		Humbeing: Humbeing{
			Name:   "lxb",
			Gender: "男",
		},
		Age: 30,
	}
	u.PrintAge()
	fmt.Println(u.Name)
}
