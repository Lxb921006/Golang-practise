package main

import (
	"fmt"
)

func main() {
	var flag *bool
	f1 := false
	flag = &f1

	c := make(chan int, 10)
	exit := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		go Write(c, exit, i)
	}

	for i := 0; i < 10; i++ {
		go Read(c, exit, flag)
	}

	for {
		if *flag {
			close(c)
			close(exit)
			break
		}
	}

	fmt.Println("finished")
}

func Write(c, e chan int, i int) {
	c <- i
}

func Read(c, e chan int, flag *bool) {
	for {
		v, k := <-c
		if !k {
			break
		}
		e <- v
		fmt.Println(v)

		if len(e) == 1000 {
			*flag = true
			break
		}
	}
}
