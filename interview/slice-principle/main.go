package main

import "fmt"

// slice原理
func main() {
	s := make([]int, 0)
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	s = append(s, s2...)
	fmt.Println(s, len(s), cap(s))

}
