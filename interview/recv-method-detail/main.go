package main

import "fmt"

// 方法接收器的类型选择: 值无法调用指针方法，指针方法集不可以让值实现接口，存储在接口的值是拿不到地址
// golang对指针可以隐式的解引用，
type Adult struct {
	Name string
}

func (a Adult) GetName() {
	fmt.Println(a.Name)
}

type Person interface {
	GetName()
}

func main() {
	var a Person = &Adult{"lxb"}
	a.GetName()
}
