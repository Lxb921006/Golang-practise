package main

import (
	"fmt"
	"time"
)

func main() {
	//当没有使用goroutine时，无缓冲通道必须要有接受才能发送，否则发成死锁
	// i := 1
	c := make(chan int, 2)
	<-c
	// c <- 1
	// c <- 2
	// c <- 3
	// <-c
	// <-c
	// close(c)
	// v, k := <-c
	// fmt.Println(v, k)
	// limit := make(chan int, 5)
	// c <- 1
	// c <- 1
	// c <- 1
	// for i := 0; i < 100; i++ {
	// 	go Recv(c, limit)
	// }

	// for {
	// 	i++
	// 	c <- i
	// }
	// <-c
	// c <- 2
	// c <- 10 //这里会收阻塞，无缓冲通道在发的时候都会阻塞直到有人接收
	// fmt.Println("send finished")
}

func Recv(c, l chan int) {
	// fmt.Println("resc succeed = ", <-c)
	for v := range c {
		l <- 1
		fmt.Printf("time = %v, val = %d\n", time.Now().Unix(), v)
		// time.Sleep(time.Second / 2)
		<-l
	}
	// c <- 1
}

func Send(c chan int) {
	c <- 1
	fmt.Println("send succeed")
}
