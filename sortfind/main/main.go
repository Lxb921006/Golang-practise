package main

import (
	"fmt"
)

func main() {
	//数组的排序和查找
	//二分查找必须是有序的
	s1 := [6]int{22, 11, 33, -2, 10, 36} //11
	s3 := arrSort(&s1, false)
	fmt.Println("s3=", s3)
	s4 := [5]string{"lxb", "lqm", "lyy"}
	filterArr01(&s4)
	s5 := []int{11, 22, 33, 44, 55, 66}
	filterArr02(s5, 0, 5, 33)
	fmt.Println("s5=", s5)
}

//冒泡排序-降序（大到小）
func arrSort(s *[6]int, s2 bool) [6]int {
	switch s2 {
	case false:
		for i := 0; i < len(s)-1; i++ {
			for t := 0; t < len(s); t++ {
				if t < len(s)-1 && s[t] < s[t+1] {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}

		}
	case true:
		for i := 0; i < len(s)-1; i++ {
			for t := 0; t < len(s); t++ {
				if t < len(s)-1 && s[t] > s[t+1] {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}
		}
	}
	return *s
}

//顺序查找
func filterArr01(s *[5]string) {
	var name string
	var out bool
	fmt.Printf("请输入名字:")
	fmt.Scanln(&name)
	for i := 0; i < len(s); i++ {
		if s[i] == name {
			fmt.Printf("%v 存在\n", s[i])
			out = true
			break
		}
	}
	if !out {
		fmt.Printf("%v 不存在\n", name)
	}
}

//二分查找
func filterArr02(n1 []int, ln2, rn3, n4 int) {
	m := (ln2 + rn3) / 2
	if ln2 > rn3 {
		fmt.Println("no")
		return
	}
	if n1[m] > n4 {
		filterArr02(n1, ln2, m-1, n4)
	} else if n1[m] < n4 {
		filterArr02(n1, m+1, rn3, n4)
	} else {
		fmt.Println("ok")
	}
}
