package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go func(i int) {
			c <- i
			fmt.Println(time.Now().UnixNano())

			time.Sleep(time.Second)
			<-c
		}(i)
	}

}
