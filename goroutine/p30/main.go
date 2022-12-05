package main

import "fmt"

func main() {
	c := make(chan struct{})
	close(c)
	// fmt.Println(<-c)
	// c <- struct{}{}
	//select-case的分支是如果都是非阻塞则随机选择一个
	select {
	// Panic if the first case is selected.
	case c <- struct{}{}:
		fmt.Println("send")
	case <-c:
		fmt.Println("recv")
	}
}
