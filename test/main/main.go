package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	res1 := Sort1()
	fmt.Println(res1)
	res2 := Sort2()
	fmt.Println(res2)
	FindPn(3)
	for i := 1; i <= 10; i++ {
		FindPn2(i, 10)
	}
}

func Sort1() []string {
	s := []string{"20", "5", "10", "50", "15", "25", "55", "35", "40", "0", "45", "30"}
	// s := []int{1, 5, 3, 7, 12, 10, 511, 40, 20, 39, 2340, 602, 103, 45, 701, 592, 958, 73, 928}
	sort.Slice(s, func(i, j int) bool {
		numA, _ := strconv.Atoi(s[i])
		numB, _ := strconv.Atoi(s[j])
		return numA < numB
	})
	return s
}

func Sort2() []string {
	s := []string{"20", "5", "10", "50", "15", "25", "55", "35", "40", "0", "45", "30"}
	// s := []int{1, 5, 3, 7, 12, 10, 511, 40, 20, 39, 2340, 602, 103, 45, 701, 592, 958, 73, 928}
	for i := 0; i < len(s)-1; i++ {
		for t := 0; t < len(s); t++ {
			if t < len(s)-1 {
				numA, _ := strconv.Atoi(s[t])
				numB, _ := strconv.Atoi(s[t+1])
				if numA > numB {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}
		}
	}
	return s
}

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
			fmt.Println("fn1是质数=", s1)
		} else {
			y = true
		}
		s1++
	}
	return true
}

func FindPn2(n, end int) {
	//大于1且除了本身不能被其他数整除就是质数
	isR := true
	for i := 1; i <= end; i++ {
		if n <= 1 {
			isR = false
			break
		}
		if n == i || i <= 1 {
			continue
		}
		if n%i == 0 {
			isR = false
			break
		}
	}
	if isR {
		fmt.Printf("%d是质数\n", n)
	}
}
