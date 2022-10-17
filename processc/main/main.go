package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// var age byte
	// fmt.Printf("请输入你的年龄:")
	// fmt.Scanln(&age)
	// if age > 18 {
	// 	fmt.Println("welcome")
	// } else {
	// 	fmt.Println("unwelcome")
	// }
	if age := 20; age > 18 {
		fmt.Println("welcome")
	} else {
		fmt.Println("unwelcome")
	}

	n := 10
	m := 10
	if n+m > 10 && n > 9 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}

	var info string
	year := 2019
	info = yearCheck(year)
	fmt.Printf("%d是%v\n", year, info)

	fmt.Println(math.Sqrt(3.0)) //平方根

	info2 := "lxb"
	index := len(info2) - 1
	switch infoInput(info2) + "aa" { // 这里可以跟表达式(表达式要加上分号)，函数
	case "lxb", "bbb": //这里可以跟多个表达式，逗号隔开
		if strings.Contains(info2, "b") {
			fmt.Printf("%v的最后一个字母是%c\n", info2, info2[index])
		}
	case "lqm": // 常量值不能重复，如这里不能再写lxb, 但是变量可以，如 bb = "lxb" ,这个是可以编译通过
		if strings.Contains(info2, "m") {
			fmt.Printf("%v的最后一个字母是%c\n", info2, info2[index])
		}
	default:
		fmt.Println("nobody")
	}

	//数据类型不一致会无法编译通过
	// var nn1 int64 = 10
	// var nn2 int64 = 10

	// switch nn1 {
	// case nn2, 20:
	// 	fmt.Println("nn")
	// }

	switch nn1 := 10; {
	case nn1 == 10:
		fmt.Println("nn1")
	case nn1 == 20:
		fmt.Println("nn2")
	default:
		fmt.Println("nn3")
	}

	// switch穿透fallthrough(只能穿透一层)，case匹配成功继续执行下面的case
	nn3 := 100
	switch nn3 {
	case 10:
		fmt.Println("nn3=", nn3)
	case 100:
		fmt.Println("nn3=", nn3)
		fallthrough
	case 200:
		fmt.Println("nn3-2=", nn3)
	default:
		fmt.Println("nn3=", 400)
	}

	//switch可以用于type-switch来判断某个interface变量中的实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型=%T", i)
	case int:
		fmt.Println("x的类型=int")
	case float64:
		fmt.Println("x的类型=float64")
	default:
		fmt.Println("未知")
	}

	//大写
	fmt.Println(strings.Title("a"))

	//无法打印i的值，因为i是局部变量
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("lxb")
	}
	t := 0
	for ; t < count; t++ {
		fmt.Println("lxb2")
	}
	t1 := 0
	for t1 < 10 {
		fmt.Println("t2")
		t1++
	}
	//死循环
	// for {
	// 	fmt.Println("t3")
	// 	time.Sleep(5 * time.Second)
	// }

	//Golang没有while，只能用for来代替
	br := 1
	for {
		if br < 10 {
			fmt.Println("1111")
		} else {
			break
		}
		br++
	}

	st := "call lxb"
	for ii := 0; ii < len(st); ii++ {
		fmt.Printf("%c", st[ii])
	}
	fmt.Println("--------------------")
	for index, val := range st {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
	fmt.Println("--------------------")
	str := "廖lxb"
	str2 := []rune(str) // 把str转成了[]rune，就是切片
	for index, val := range str2 {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
	fmt.Println("---------三角形-----------")
	for t := 1; t <= 5; t++ {
		fmt.Println()
		for ttt := 1; ttt <= 5-t; ttt++ {
			fmt.Printf(" ")
		}
		for tt := 1; tt <= 2*(t-1)+1; tt++ {
			if tt == 1 || tt == 2*(t-1)+1 || t == 5 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
	}
	fmt.Println("----------九九乘法----------")
	for i := 1; i <= 9; i++ {
		for t := 1; t <= i; t++ {
			fmt.Printf("%d * %d = %d \t", t, i, i*t)
		}
		fmt.Println()
	}
	fmt.Println("----------随机数----------")
	rr1 := myRand("go")
	fmt.Println("rr1的随机数=", rr1)
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	// var password string
	// cc := 1
	// cc2 := 3
	// for {
	// 	if cc <= 3 {
	// 		fmt.Printf("请输入密码:")
	// 		fmt.Scanln(&password)
	// 		if password != "abc" {
	// 			cc2--
	// 			if cc2 == 0 {
	// 				fmt.Printf("只有三次机会, bye\n")
	// 				break
	// 			}
	// 			fmt.Printf("密码错误, 还有%d机会\n", cc2)
	// 		} else {
	// 			fmt.Println("welcome")
	// 			break
	// 		}
	// 	}
	// 	cc++
	// }
	fmt.Println("----------goto----------")
	bb := 10
	if bb > 2 {
		goto label1 //一般跟if语句配合使用
	}
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
label1:
	fmt.Println(3)
	fmt.Println(5)
	fmt.Println("----------return----------")
	for b := 1; b < 5; {
		if b == 3 {
			return
		}
		fmt.Println(b)
		b++
	}
}

//------------------函数------------------------
func infoInput(s string) string {
	return s
}

func yearCheck(i int) string {
	var info string
	if (i%4 == 0 && i%100 != 0) || i%4 == 0 {
		info = "闰年"
	} else {
		info = "平年"
	}
	return info
}

func myRand(s string) int {
	var nn int
	if s == "go" {
		r1 := rand.NewSource(time.Now().Unix())
		r2 := rand.New(r1)
		nn = r2.Intn(100)
	} else {
		nn = 100
	}
	return nn
}
