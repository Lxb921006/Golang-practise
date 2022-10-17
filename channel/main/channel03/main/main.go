package main

import "fmt"

func main() {
	fmt.Println("----------------管道的只读,写的最佳实践------------------")
	//管道默认是可以读写, select解决从管道获取数据的阻塞问题
	chan2 := make(chan int, 3)
	//管道只写
	Wchan(chan2)
	//管道只读
	//go Rchan(chan2, &exFlag)

	for {
		select {
		case v := <-chan2:
			fmt.Println(v)
		default:
			fmt.Println("finished....")
			return
		}
	}
}

func Wchan(c chan<- int) {
	for i := 1; i <= 3; i++ {
		c <- i
	}
}

// func Rchan(c <-chan int) {
// 	for i := 1; i <= 3; i++ {
// 		n := <-c
// 		fmt.Println(n)
// 	}
// }
