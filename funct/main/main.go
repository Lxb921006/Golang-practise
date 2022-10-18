package main

//如果调用其他的包要把，引用别的包的名字需声明为package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/Lxb921006/Golang-practise/funct/utils" //给包起个别名
)

func main() {
	// fmt.Println("res=", res)
	// fmt.Println("name=", util.Name)
	// test1(4)
	// test2(4)
	// res2 := test3(3)
	// fmt.Println("res2=", res2)
	res := util.CountNum(10, 20)
	res3 := test4(4)
	fmt.Println("res3=", res3, res)
	res4 := test5(1)
	fmt.Println("res4=", res4)
	res5 := 10
	//可以将函数赋值给一个变量，该变量就是一个函数类型的变量
	res7 := test6
	fmt.Println("main res7=", res7(&res5))

	res8 := func(n int) int {
		return n
	}
	fmt.Println("res8=", res8(8))

	//将函数当做参数传入
	res9 := test7(test8, 100)
	fmt.Println("res9=", res9)

	res10 := test72(test8, 10000)
	fmt.Println("res10=", res10)

	//支持对函数返回值命名
	res11, res12 := sumSub(10, 9)
	fmt.Printf("res11=%d, res12=%d\n", res11, res12)

	//支持可变参数
	res13 := test9(10, 10)
	fmt.Println("res13=", res13)

	res14 := test10(10, 20)
	fmt.Println("res14=", res14)

	//自定义数据类型的使用，但是如下定义在golang里边是两个不同类型，即，cus 不等价于 int
	type cus int
	var num2 int
	var num1 cus = 10

	//num2 = num1 这个是报错，只能显式转类型
	num2 = int(num1)
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	//例子
	res15 := test11(1.2, 1.8)
	fmt.Println("res15=", res15)

	res16, res17 := 11, 22
	fmt.Printf("res16没变之前=%d, res17没变之前=%d\n", res16, res17)
	res18, res19 := test12(&res16, &res17)
	fmt.Printf("res16变化之后=%d, res17变化之后=%d\n", res18, res19)

	//函数金字塔案例
	test13(5)

	//函数99乘法表
	test15(5)

	//字符串常用函数,golang的编码统一为utf-8，ascii字符（字母和数字），占用一个字节，中文占用3个字节
	//str1 := "lxb"
	//len(str)
	//[]rune()
	str1 := "lxb中国人"
	str2 := []byte(str1) //切片,遍历有中文字符
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2=%c\n", str2[i])
	}
	//strconv.Itoa整数转字符串
	str3 := strconv.Itoa(100)
	fmt.Println("str3=", str3)
	//strconv.Atoa字符串转整数
	str4, _ := strconv.Atoi("100")
	fmt.Println("str4=", str4)
	//字符串转成[]byte()切片,
	var str5 = []byte("lxb lqm")
	fmt.Println("str5=", str5) //输出ascii码
	//[]byte转成字符串
	str6 := string([]byte{97, 98})
	fmt.Println("str6=", str6)
	//十进制转2，8，16进制
	str7 := strconv.FormatInt(123, 2)
	fmt.Println("str7=", str7)
	//查找目标字符串中是否包含某个子串
	str8 := "lxblqm"
	fmt.Println("str8=", strings.Contains(str8, "lxb"))
	//统计一个字符串里边有几个指定的子串
	str9 := "lxblqm"
	fmt.Println("str9=", strings.Count(str9, "l"))
	//不区分大小写的字符串比较（==是区分大小写），
	fmt.Println("str10=", strings.EqualFold("abc", "Abc"))
	//返回子串在字符串第一次出现的index值，如果没有返回-1
	fmt.Println("str11=", strings.Index("cccc_abc", "abc"))
	fmt.Println("str11=", strings.LastIndex("cccc_abcabc", "abc"))
	//指定的子串替换成另一个子串, Replace的最后一个参数如果等于-1则表示全部替换
	fmt.Println("str12=", strings.ReplaceAll("lxb lxb lqm", "lxb", "lll"))
	fmt.Println("str13=", strings.Replace("lxb lxb lqm", "lxb", "lll", 1))
	//拆分字符串
	fmt.Println("str14=", strings.Split("lxb,lqm", ","))
	str14 := strings.SplitN("lxb,lqm,lxb", ",", 2)
	fmt.Println("str14=", len(str14))
	fmt.Println("str14=", str14[len(str14)-1])
	fmt.Println("str14=", str14[0:])
	//字母大小写转换
	fmt.Println("str15=", strings.ToLower("LXB"))
	fmt.Println("str15=", strings.ToUpper("lqm"))
	//将字符串左右两边指定的字符去掉
	fmt.Println("str17=", strings.Trim("(lqm)", "()"))
	fmt.Println("str17=", strings.TrimSpace(" ( lqm) "))
	fmt.Println("str17=", strings.TrimLeft(strings.TrimSpace(" ( lqm) "), "("))
}

// func test1(n int) {
// 	if n > 2 {
// 		n--
// 		test1(n)
// 	}
// 	fmt.Println("test1 n = ", n)
// }

// func test2(n int) {
// 	if n > 2 {
// 		n-- //必须向递归条件逼近，否则变死归
// 		test2(n)
// 	} else {
// 		fmt.Println("test2 n = ", n)
// 	}
// }

// //斐波那契数
// func test3(n int) int {
// 	if n == 1 || n == 2 {
// 		return 1
// 	}
// 	return test3(n-1) + test3(n-2)
//
// }

func test4(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*test4(n-1) + 1
	}
}

func test5(n int) int {
	if n == 10 {
		return 1
	} else {
		return (test5(n+1) + 1) * 2
	}
}

func test6(n *int) int {
	fmt.Println("test6 *n=", *n)
	*n += 10
	return *n
}

//将函数当做参数传入
func test7(funct func(int) int, n1 int) int {
	return funct(n1)
}

type testFunct func(n int) int //自定义函数的数据类型类型，这个很有用

func test72(funct testFunct, n1 int) int {
	return funct(n1)
}

//支持对函数返回值命名, 这样好处就是不用在return 返回值
func sumSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//支持可变参数，即支持多个参数,其实args就是切片, args...必须这种格式，且必须放在参数的最后一个
func test9(n int, args ...int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func test8(n int) int {
	return n
}

func test10(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func test11(n1, n2 float32) float32 {
	return n1 + n2
}

func test12(n1, n2 *int) (n3, n4 int) {
	n3, n4 = *n2, *n1
	return
}

//案例
func test13(n int) {
	for i := 1; i <= n; i++ {
		for g := 1; g <= n-i; g++ {
			fmt.Printf(" ")
		}
		for t := 1; t <= 2*(i-1)+1; t++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

func test15(n int) {
	for i := 1; i <= n; i++ {
		for t := 1; t <= i; t++ {
			fmt.Printf("%d * %d = %d \t", t, i, i*t)
		}
		fmt.Println()
	}

}
