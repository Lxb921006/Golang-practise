package main

import "fmt"

func main() {
	num := 0

	for {
		fmt.Println("输入:")
		fmt.Scanln(&num)
		num++
		fmt.Println("num = ", num)
	}
}
