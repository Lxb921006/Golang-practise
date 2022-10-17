package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	s1[0] = 100
	fmt.Println("s2 = ", s2)
	fmt.Println("s1 = ", s1)
}
