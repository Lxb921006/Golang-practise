package main

import (
	"sync"
	_ "time"
)

var (
	wg  sync.WaitGroup
	wr  = make(chan int, 4) //这里用管道作用是限制go协程数量，再没有被消费的情况下每次管道内只能有4个队列超过就阻塞
	rwr = make(chan int)
)

func main() {
	//练习3,求质数
	for i := 1; i <= 80000; i++ {
		wr <- i
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// time.Sleep(time.Second)
			FindPn2(i, 80000)
			<-wr
			<-rwr
		}(i)
	}
	wg.Wait()
}

func FindPn2(n, end int) {
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
	}
}
