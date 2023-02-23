package main

import "fmt"

// slice原理
func main() {
	s := make([]int, 0)
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	s = append(s, s2...)
	fmt.Println(s, len(s), cap(s))

	s3 := []int{22, 33, 44, 55, 66, 77, 88, 99, 11}
	s4 := s3[5:]
	s5 := s3[:5]

	fmt.Println(s4)
	fmt.Println(s5)

}
