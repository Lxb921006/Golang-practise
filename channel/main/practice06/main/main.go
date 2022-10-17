package main

import (
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	wr := make(chan int, 100)
	rwr := make(chan int, 100)
	go Write(wr)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go Read(wr, rwr)
	}
	wg.Wait()
}

func Write(wr chan int) {
	for i := 1; i <= 80000; i++ {
		wr <- i
	}
}

func Read(wr chan int, rwr chan int) {
	defer wg.Done()
	for {
		if len(wr) == 0 {
			break
		}
		d := <-wr
		FindPn2(d, 80000, rwr)
	}
}

func FindPn2(n, end int, rwr chan int) {
	//大于1且除了本身不能被其他数整除就是质数
	isR := true
	for i := 1; i <= end; i++ {
		if n <= 1 {
			isR = false
			break
		}
		if n == i || i <= 1 {
			continue
		}
		if n%i == 0 {
			isR = false
			break
		}
	}
	if isR {
		rwr <- n
		<-rwr
	}
}
