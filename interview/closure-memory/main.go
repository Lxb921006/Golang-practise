package main

import (
	"fmt"
	"time"
)

// 为什么闭包会具有记忆效应-需要进行逃逸分析，nn局部变量从栈区释放但会发生逃逸到堆区，属于堆里边的变量，因此闭包也具有记忆效应
// 主要是因为闭包引用了该局部变量的地址

func main() {
	c := add()
	fmt.Println(c())
	time.Sleep(time.Second * 10)
	fmt.Println(c())
	fmt.Println(c())
}

func add() func() int {
	var nn int
	return func() int {
		nn++
		return nn
	}
}
