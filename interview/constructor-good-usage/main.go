package main

import "fmt"

// 优雅的构造函数写法
type Person struct {
	Name string
	Tel  int
}

type Option func(*Person)

func NewPerson_(options ...Option) *Person {
	p := &Person{}
	for _, option := range options {
		option(p)
	}
	fmt.Println(p)
	return p
}

func CallName(name string) Option {
	return func(p *Person) {
		p.Name = name
	}
}

func CallAge(age int) Option {
	return func(p *Person) {
		p.Tel = age
	}
}

func main() {

	_ = NewPerson_(
		CallName("lxb"),
		CallAge(31),
	)
}
