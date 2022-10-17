package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func OnlyOnce() {
	fmt.Println("once")
}

// 单例模式演示-dev
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go func(t int) {
			once.Do(OnlyOnce)
			c <- t
		}(i)
	}

	for i := 0; i < 5; i++ {
		d := <-c
		fmt.Println(d)
	}
}
