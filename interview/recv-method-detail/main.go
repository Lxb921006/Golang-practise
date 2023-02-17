package main

import "fmt"

// 方法接收器的类型选择
type Adult struct {
	Name string
}

func (a *Adult) GetName() {
	fmt.Println(a.Name)
}

type Person interface {
	GetName()
}

func main() {
	var a Person = &Adult{"lxb"}
	a.GetName()
}
