package main

import "fmt"

func main() {
	//二维数组
	// 定义二维数组
	a1 := [4][2]int{{11, 22}, {33, 44}, {55, 66}, {77, 88}}
	fmt.Println("a1=", a1)
	for i := 0; i < len(a1); i++ {
		for t := 0; t < len(a1[i]); t++ {
			fmt.Printf("%d ", a1[i][t])
		}
		fmt.Println("")
	}
}
