package main

import (
	"fmt"
	"sync"
	"time"
)

//限制goroutine数量
var (
	wg sync.WaitGroup
	c  = make(chan int, 20)
)

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(t int) {
			c <- t
			fmt.Println("date= ", time.Now().Unix())
			time.Sleep(time.Second)
			<-c
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished!!!")
}
