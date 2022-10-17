package main

import "fmt"

//工厂模式
type Animal interface {
	Flay()
	Swim()
}

type Fish struct {
}

func (Fish) Flay() {
	fmt.Println("fish flay")
}

func (Fish) Swim() {
	fmt.Println("fish flay")
}

type Bird struct {
}

func (Bird) Flay() {
	fmt.Println("Bird flay")
}

func (Bird) Swim() {
	fmt.Println("Bird flay")
}

//返回Animal接口,只要实现了Animal接口的方法,都可以调用-
func NewAnimal(i int) Animal {
	switch i {
	case 1:
		return &Bird{}
	case 2:
		return &Fish{}
	default:
		return nil
	}
}

func main() {
	bird := NewAnimal(1)
	bird.Flay()

	fish := NewAnimal(2)
	fish.Flay()
}
