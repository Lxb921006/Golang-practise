package main

import "fmt"

//继承跟接口的关系
type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Printf("%v会爬树\n", m.Name)
}

type Monkey01 struct {
	Monkey
}

func (m *Monkey01) Flying() { //虽然没有调用接口，但是它确实实现了接口的方法，因为接口的方法定义了Flying()
	fmt.Printf("%v会飞\n", m.Name)
}

func (m *Monkey01) Swiming() {
	fmt.Printf("%v会游泳\n", m.Name)
}

//声明接口
type Bird interface {
	Flying()
}

//声明接口
type Fish interface {
	Swiming()
}

func main() {
	m := Monkey01{}
	(&m).Name = "猴子"
	// (&m).Climbing()
	// m.Flying()
	// m.Swiming()
	var m1 = &m
	m1.Flying()
}
