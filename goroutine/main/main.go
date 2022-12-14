package main

import (
	"fmt"
	"strconv"
	"time"
)

//注意事项
//1.MPG模型中的一个M对应内核空间里的一个内核线程,相当于内核线程在Go进程中的映射
//2.每一个M都会以一个内核线程绑定,M和P之间也是一对一的关系.P和G的关系则是一对多
//3.运行过程中,M线程之间的对应关系不会变化,在M的生命周期内,它只会与一个内核线程
//绑定,而M,P,G之间的关系都是可变的
//4.只有M和P的组合才能为G提供有效的运行环境,多个可执行的G会排成一个队列挂在某个
//P上等待调度和执行
//5.M的创建一般是因为没有足够的M来和P组合为G提供运行环境,很多时候M的数量可能会
//比P要多,单个G进程中,P的最大数量决定了程序的并发规模,由程序决定,修改GOMAXPROCS
//或者调用函数runtime来设定P的最大值
//6.M和P会适时的组合和断开来保证P中的待执行G队列能够得到及试运行,如果有G1因为网络
//等阻塞了M,P就会携带剩余的G投入到其他的M1中,这个M1可能是新创建的,也可能是从调度器
//空闲的M列表中获取的,取决于此时调度器空闲M列表中是否存在M,从而避免M的过多创建

//重点
//7.当挂起的M对应的内核线程被唤醒时,M将会尝试为G1捕获一个P的上下文,可能从调度器的空闲
//队列中获取,如果获取不成功,M就会被G1放入到调度器的可执行G队列中,等待其他P的查找,为了保证
//G的均衡执行,非空闲的P会运行完自身的可执行G队列中,会周期性的从调度器的可执行G队列中获取
//可执行的G,甚至从其他的P的可执行G队列中掠夺G

// ——————————————————————————————————————————————————————————————————————
//8.Golang的并发真理:不要以共享内存的方式来通信,而是通过通信来共享内存
// ———————————————————————————————————————————————————————————————————————

//9.Golang的并发看上去是抢占式的,实际上协程并非抢占式,也就是当协程遇到
//阻塞时,是当前的被阻塞的协程主动交出控制权,而不是调度器强制切换

func main() {
	//goroutine协程的使用
	//案列，统计1-20000数字中的素数
	//Go协程特点：有独立的栈空间，共享程序堆空间，调度由用户控制，是轻量级的线程
	//Go的主线程，也可以理解为进程，可以起多个协程

	FindPn(1)
	fmt.Println("---------------协程入门案列--------------")
	//主线程main()完成了，协程不管有没有完成都退出程序
	//主线程是一个物理线程，直接作用在cpu上。是重量级的，非常耗费cpu资源，由操作系统控制
	//协程从主线程开启，是轻量级线程，是逻辑态，对资源消耗小
	//Golang的协程是重要的特点，可以轻松开启上万个协程

	go test01() //开启一个协程
	time.Sleep(time.Second * 2)
	for i := 1; i <= 5; i++ {
		fmt.Println("hello go_" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println("---------------协程调度模型MPG状态1(静态)--------------")
	//M:操作系统的主线程（是物理协程），P:协程执行需要上下文（运行环境，cpu，内存等等），G:协程

	fmt.Println("---------------协程调度模型MPG状态2(动态)--------------")

	fmt.Println("---------------Golang运行的cpu数量设置--------------")
	//在1.8之前,需要设置,1.8之后就不需要了,默认会作用在多核上

}

func test01() {
	for i := 1; i <= 5; i++ {
		fmt.Println("hello_" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//传统的方法统计素数
func FindPn(n int) bool {
	s1 := 2
	for {
		y := true
		s2 := 2
		if s1 == n {
			break
		}
		for {
			if s2 == n {
				break
			}
			if s1 == s2 {
				s2++
				continue
			}
			if s1%s2 == 0 {
				y = false
				break
			}
			s2++
		}
		if y {
			fmt.Println("是质数=", s1)
		} else {
			y = true
		}
		s1++
	}
	return true
}
