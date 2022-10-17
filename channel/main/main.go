package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1 = make(map[int]int, 10)
	//声明一个全局的互斥锁,低级程序
	lock sync.Mutex
)

func main() {
	//全局互斥锁的使用,及管道的引入
	for n := 1; n <= 20; n++ { //启动20个协程
		go test01(n)
	}

	time.Sleep(time.Second * 5) //全局互斥锁虽然可以解决资源竞争问题,但是这里的阻塞直到所有协程执行完到底需要多少秒合适主线程并不知道,所以需要用到管道:channel
	//如果不设置阻塞，有些协程可能都还没有执行完就因为主线程退出而退出了

	//直接输出m1会报错fatal error: concurrent map writes，这里因为上面的协程没有阻塞的情况下，由于资源共享问题直接遍历m1会报错，不像老版本之前输出空
	//查看资源是否存在竞争问题可以用：go build -race main.go
	//在Golang协程里,只要是在同一时间(读写)共享资源,(该案例的共享资源就是全局变量m1)时在没有做处理的情况下,都会出现资源竞争问题
	lock.Lock() //这里也要加锁是因为，主线程并不知道10秒内可以跑完20个协程，底层还会出现资源竞争问题
	for i, v := range m1 {
		fmt.Println(i, v)
	}
	lock.Unlock()
}

func test01(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res += i //出现v=0是因为阶乘数值太大了，已经越界出现0, 改成累加
	}
	lock.Lock()
	m1[n] = res
	lock.Unlock()
}
