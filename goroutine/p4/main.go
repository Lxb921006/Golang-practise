package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//限制goroutine数量
var (
	wg sync.WaitGroup
	c  = make(chan int, 20)
)

func main() {
	for i := 0; i < 10000; i++ {
		c <- 1
		wg.Add(1)
		go recv(c)
	}
	wg.Wait()
	fmt.Println("finished!!!")
}

func recv(c chan int) {

	defer wg.Done()

	fmt.Printf("date= %v, gn = %d\n", time.Now().Unix(), runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
	<-c

}
