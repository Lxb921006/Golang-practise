package main

import "fmt"

func main() {
	c := make(chan int, 5)
	// close(c)
	// panic:all goroutines are asleep - deadlock
	// <-c

	c <- 1
	fmt.Println("发送完毕")
}
