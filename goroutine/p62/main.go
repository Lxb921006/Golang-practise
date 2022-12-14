package main

import (
	"log"
	"time"
)

func main() {
	c1 := make(chan int)
	// c2 := make(chan int)

	func1 := func() {
		select {
		case c1 <- 1:
		default:
		}
	}

	func2 := func() {
		for {
			select {
			case <-c1:
				return
			default:
				log.Print("adada")
			}
		}
	}

	for i := 0; i < 3; i++ {
		func1()
	}

	go func2()

	time.Sleep(time.Second * 10)

}
