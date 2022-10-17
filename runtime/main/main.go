package main

import (
	"fmt"
	"runtime"
)

func main() {
	//设置golang运行的需要的cpu数量

	cpuN := runtime.NumCPU()

	runtime.GOMAXPROCS(cpuN) //在1.8之前,需要设置,1.8之后就不需要了,默认会作用在多核上

	fmt.Println(cpuN)

	r := 2
	r *= 3
	fmt.Println(r)

	n := 1

	for i := 1; i <= 100; i++ {
		n *= i
	}
	fmt.Println(n)
}
