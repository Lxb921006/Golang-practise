package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup //主线程等待所有协程都执行完毕再退出
)

func main() {
	//管道跟goroutine的综合练习
	wr := make(chan int, 10)
	go Write(wr)
	go Read(wr)
	wg.Add(1) //内部会维护着一个计数器为1,每Done()一次就减一,直到为0就释放主线程,要放在wait()前面,否则可能有些线程没执行完因为主线程的退出就退出了
	wg.Wait() //阻塞主线程,直到计数器的值减为0

}

func Write(wr chan int) {
	for i := 1; i <= 50; i++ {
		wr <- i
		fmt.Println("写=", i)
	}
	close(wr)
}

func Read(wr chan int) {
	defer wg.Done() //这里会每次让计数器减一
	for {
		v, ok := <-wr
		if !ok {
			break
		}
		fmt.Println("读=", v)
	}
}
