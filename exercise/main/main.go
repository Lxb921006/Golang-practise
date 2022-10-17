package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	res := myRang()
	fmt.Println("res=", res)
	date, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-03-37 17:36:17", time.Local)
	fmt.Println("date2=", date.Unix())
	fmt.Println("date=", time.Unix(time.Now().Unix(), 0).Format("2006-01-12 15:04:15"))
	fmt.Println("unix=", time.Now().Unix())
	fmt.Println("--------------------数组练习-------------------")
	var letter [26]byte
	for i := 0; i < len(letter); i++ {
		letter[i] = 'A' + byte(i)
	}
	f2 := [4]int{10, 20, 5, -2}
	f3 := 0
	i1 := 0
	//最大值
	for i := 0; i < len(f2); i++ {
		if f2[i] > f3 {
			i1 = i
			f3 = f2[i]
		}
	}
	fmt.Printf("f2最大的元素=%d, 下标=%d\n", f3, i1)
	//最小值
	for i := 0; i < len(f2); i++ {
		if f2[i] < f3 {
			i1 = i
			f3 = f2[i]
		}
	}
	fmt.Printf("f2最小的元素=%d, 下标=%d\n", f3, i1)
	// fmt.Println("f2最大的元素=", math.Max(1.0, 2.0))
	//细节注意，当int数值跟int数值相除，在golang里边如果有小数点，则会被遗弃，需要转数据类型如fload64(15)/fload64(2)
	var f4 [4]int
	for i := len(f2) - 1; i >= 0; i-- {
		f4[len(f2)-1-i] = f2[i]
	}
	fmt.Println("f4=", f4)
	fmt.Println("f2=", f2)
	s1 := sliceE1(6)
	fmt.Println("s1=", s1)
	fmt.Println("--------------------数组练习-------------------")
	avgCount()
	turnNum()
	s2 := [6]int{22, 6, 23, 57, 1, 6}
	s3 := arrSort(&s2, true)
	fmt.Println("s3=", s3)
	tdExer()
	tdExer02()
	filterWord("AA")
	score()
	fmt.Println("--------------------map练习-------------------")
	m1 := make(map[string]map[string]string)
	// m1["user01"] = map[string]string{"name": "lxb", "password": "0000000"}
	fmt.Println("m1-1=", m1)
	mapExerc(m1, "user01")
	fmt.Println("m1-2=", m1)
	mapExerc(m1, "user01")
	fmt.Println("m1-3=", m1)
	fmt.Println("--------------------函数练习-------------------")
	game()
}

func myRang() int {
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)
	nn := r2.Intn(100)
	return nn
}

func game() {
	n := 1
	var num int
	var out bool
	for {
		if num == 0 {
			break
		}
		if n == 11 || out {
			break
		}
		res := myRang()
		fmt.Println("num=", num)
		fmt.Printf("请输入1-100数字,有十次机会(%d):", res)
		fmt.Scanln(&num)
		switch n {
		case 1:
			if res == num {
				fmt.Println("牛逼")
				out = true
			} else {
				fmt.Printf("第%d次猜答案错误请继续\n", n)
			}
		case 2, 3:
			if res == num {
				fmt.Println("可以喔")
				out = true
			} else {
				fmt.Printf("第%d次猜答案错误请继续\n", n)
			}
		case 4, 5, 6, 7, 8, 9:
			if res == num {
				fmt.Println("一般般")
				out = true
			} else {
				fmt.Printf("第%d次猜答案错误请继续\n", n)
			}
		case 10:
			if res == num {
				fmt.Println("终于猜对了")
				out = true
			} else {
				fmt.Printf("好菜,十次一次都没猜对, byebye")
			}
		default:
			out = true
		}
		n++
	}
}

//切片练习
func sliceE2(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return sliceE2(n-2) + sliceE2(n-1)
}

func sliceE1(n int) []uint64 {
	s1 := make([]uint64, n)
	for i := 0; i < n; i++ {
		s1[i] = sliceE2(i + 1)
	}
	return s1
}

func avgCount() {
	s := [3][5]int{{10, 11, 12, 13, 14}, {15, 10, 12, 14, 16}, {14, 11, 17, 19, 13}}
	all := .0
	for i := 0; i < len(s); i++ {
		one := 0
		one_avg := .0
		for t := 0; t < len(s[i]); t++ {
			one += s[i][t]
		}
		one_avg = float64(one) / float64(len(s[i]))
		all += one_avg
		fmt.Printf("s[%d] avg = %.2f\n", i, one_avg)
	}
	fmt.Printf("all avg = %.2f\n", all/float64(len(s)))
}

//反转
func turnNum() {
	s1 := [5]int{11, 5, 16, 9, 3}
	for i := 0; i < len(s1); i++ {
		if (len(s1)-1)/2 >= i {
			tmp := s1[i]
			s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], tmp
		}
	}
	fmt.Println("s1=", s1)
}

//二维练习
func tdExer() {
	s1 := [3][4]int{{11, 22, 33, 44}, {55, 55, 55, 55}, {66, 66, 66, 66}}
	for i := 0; i < len(s1); i++ {
		for t := 0; t < len(s1[i]); t++ {
			if i == 0 || i == len(s1)-1 {
				s1[i][t] = 0
			} else {
				if t == 0 {
					s1[i][t] = 0
				}
				if t == len(s1[i])-1 {
					s1[i][t] = 0
				}
			}
			fmt.Printf("%d ", s1[i][t])
		}
		fmt.Println()
	}
}

func tdExer02() {
	s1 := [3][4]int{{11, 22, 33, 44}, {55, 55, 55, 55}, {66, 66, 66, 66}}
	for i := 0; i < len(s1); i++ {
		if (len(s1)-1)/2 >= i {
			tmp := s1[i]
			s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], tmp
		}
	}
	fmt.Println("s1=", s1)
}

func filterWord(s string) {
	s1 := [5]string{"AA", "la", "jt", "AA"}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s {
			fmt.Printf("s1[%d]=%v\n", i, s1[i])
		}
	}
}

func score() {
	total := .0
	goodSocre := .0
	goodReferee := 0
	s1 := [8]float64{70.5, 81.6, 90, 83.5, 94.6, 80, 86, 90.2}
	s2 := s1[:]
	max, i1 := findNum(s1[:], true)
	min, i2 := findNum(s1[:], false)
	for i := 0; i < len(s1); i++ {
		if s1[i] == max || s1[i] == min {
			s1[i] = .0
		}
		total += s1[i]
	}
	avg, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", total/float64(6)), 64)
	s3 := append(s2, avg)
	s4 := sliceSort(s3, true)
	for i := 0; i < len(s4); i++ {
		if s4[i] == avg {
			if i == 0 {
				goodSocre = s4[i+1]
				goodReferee = i + 1
			} else {
				ln := s4[i] - s4[i-1]
				rn := s4[i] - s4[i+1]
				if rn > ln {
					goodSocre = s4[i-1]
					goodReferee = i - 1
				} else {
					goodSocre = s4[i+1]
					goodReferee = i + 1
				}
			}
			break
		}
	}
	fmt.Printf("s4=%v\n", s4)
	fmt.Printf("max=%v, index=%d\n", max, i1)
	fmt.Printf("min=%v, index=%d\n", min, i2)
	fmt.Printf("avg=%.2f\n", avg)
	fmt.Printf("最佳评委=%d,分数=%v\n", goodReferee, goodSocre)
	fmt.Printf("最差评委=%d,分数=%v\n", len(s4)-1, s4[len(s4)-1])
}

func findNum(s []float64, s2 bool) (float64, int) {
	num := .0
	i1 := 0
	switch s2 {
	case true:
		for i := 0; i < len(s); i++ {
			if s[i] > num {
				i1 = i
				num = s[i]
			}
		}
	case false:
		num = 10000.0
		for i := 0; i < len(s); i++ {
			if s[i] < num {
				i1 = i
				num = s[i]
			}
		}
	}
	return num, i1
}

func arrSort(s *[6]int, s2 bool) [6]int {
	switch s2 {
	case false: //降序，大到小
		for i := 0; i < len(s)-1; i++ {
			for t := 0; t < len(s); t++ {
				if t < len(s)-1 && s[t] < s[t+1] {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}

		}
	case true: //升序，小到大
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

func sliceSort(s []float64, s2 bool) []float64 {
	switch s2 {
	case false: //降序，大到小
		for i := 0; i < len(s)-1; i++ {
			for t := 0; t < len(s); t++ {
				if t < len(s)-1 && s[t] < s[t+1] {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}
		}
	case true: //升序，小到大
		for i := 0; i < len(s)-1; i++ {
			for t := 0; t < len(s); t++ {
				if t < len(s)-1 && s[t] > s[t+1] {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}
		}
	}
	return s
}

//map练习
type mapType map[string]map[string]string

func mapExerc(u mapType, n string) {
	// _, res := u[n]
	// u[n] = make(map[string]string)
	// if res {
	// 	u[n]["password"] = "8888888"
	// } else {
	// 	u[n]["password"] = "6666666"
	// }
	if u[n] != nil {
		u[n]["password"] = "8888888"
	} else {
		u[n] = make(map[string]string, 10)
		u[n]["password"] = "6666666"
	}
}
