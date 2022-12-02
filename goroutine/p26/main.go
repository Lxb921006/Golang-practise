package main

import "fmt"

func main() {
	cc := make(chan int)

	recv := func() {
		for {
			fmt.Println(<-cc, "退不出来")
			cc <- 1
		}
	}

	go recv()
	go recv()

	cc <- 1

	var c chan bool // nil
	<-c             // blocking here for ever

	fmt.Println("end")
}
