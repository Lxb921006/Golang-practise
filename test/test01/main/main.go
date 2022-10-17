package main

import (
	"fmt"
	"time"
)

func main() {

	num := make(chan int) //无缓冲
	letter := make(chan int)
	result := make(chan int, 2)

	go Letter(letter)
	go Num(num)
	go GetNum(num, result)
	go GetLetter(letter, result)

	//阻塞主线程直到所有协程跑完
	for {
		if len(result) == 2 {
			close(result)
			break
		}
	}

	fmt.Println("finished")

}

func Num(num chan int) {
	for i := 0; i < 10; i++ {
		num <- i
	}
	close(num)
}

func Letter(letter chan int) {
	for i := 97; i < 107; i++ {
		letter <- i
	}
	close(letter)
}

func GetNum(num, result chan int) {
	for {
		k, v := <-num
		if !v {
			result <- 1 //读完所有数字就往主线程的result写入一个标记
			break
		}
		fmt.Println(k)
		time.Sleep(time.Second)
	}
}

func GetLetter(letter, result chan int) {
	for {
		k, v := <-letter
		if !v {
			result <- 1 //读完所有字母就往主线程的result写入一个标记
			break
		}
		fmt.Printf("%c\n", k)
		time.Sleep(time.Second)
	}

}
