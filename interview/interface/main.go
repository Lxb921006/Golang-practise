package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) GetName() {
	fmt.Println("name =", p.Name)
}

type Man struct {
	Person
}

type Kind interface {
	GetName()
}

func GetName(p Kind) {
	p.GetName()
}

func main() {
	m := Man{
		Person{Name: "lxb"},
	}

	GetName(m)
}
