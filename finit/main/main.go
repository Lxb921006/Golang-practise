package main

import (
	"fmt"
	_ "os"
	"strconv"
	"strings"

	util "github.com/Lxb921006/Golang-practise/finit/utils"
) //给包起个别名

//每一个源文件都可以包含一个init函数，init函数会在main函数执行前，被go运行框架调用，也就是说init会在main函数前被调用
//可以完成初始化工作
//1.有全局变量，则顺序是全局变量->init()->main()

var test_init = test(10)

var (
	F3 = func(n int) int {
		return n
	}
)

func test(n int) int {
	fmt.Println("test()")
	return n
}

func init() {
	fmt.Println("init()")
}

func add() func(int) int {
	//以下部分就是闭包
	var n int = 10 //这里的n只会初始化一次，相当于当第一次调用时，n=10,然后第二次调用时，n+=10
	var n1 = "lxb"
	return func(i int) int {
		n += i
		n1 += strconv.FormatInt(int64(i), 10)
		fmt.Printf("n1=%v\n", n1)
		return n
	}
}

//返回指定的文件类型
func fileType(str1 string) func(string) string {
	return func(str2 string) string {
		if strings.HasSuffix(str2, str1) {
			return str2
		}
		return str2 + str1
	}
}

//defer的基本使用：在创建资源时，如数据库连接，文件句柄，锁等，为了在函数执行完毕后及时释放资源，所以有了defer（延时机制）
func useDefer(n int) int {
	defer fmt.Println("defer n1=", n) //会将这里的defer后面的语句加入到独立的栈中，暂时不会执行，当函数执行完毕后，按照先入后出的方式出栈，执行
	defer fmt.Println("defer n2=", n)
	n++ //如果n=1,这里的n++=2后的n值并不会影响上面defer语句后面打印的n值，是在不同的栈中，所以defer后面的n还是1
	return n
}

//defer例子
// func fileOperate(f string) string {
// 	file, _ := os.Open("main.go")
// 	defer file.Close()
// 	content := file.Read()
// 	return content
// }

func main() {
	fmt.Printf("name=%v, age=%d\n", util.Name, util.Age)
	fmt.Println("main(), test_init=", test_init)

	//匿名函数
	//方式1
	f1 := func(n1 int) int {
		return n1
	}(10)
	fmt.Println("f1=", f1)

	//方式二，匿名函数赋值给变量调用
	f2 := func(n1 int) int {
		return n1
	}
	fmt.Println("f2=", f2(20))
	// fmt.Println("f2-1=", f2(30))
	// fmt.Printf("f2的类型=%T", f2)

	//方式三全局变量
	fmt.Println("F3=", F3(30))

	//闭包：是一个函数和与其相关的引用环境组合的一个整体
	f4 := add()
	fmt.Println("f4=", f4(1))
	fmt.Println("f4-2=", f4(2))
	//判断文件是否有后缀
	res := strings.HasSuffix("a.txt", ".txt")
	fmt.Println("res=", res)

	res1 := fileType(".jpg") //很像python类的里边一个方法，实例化类，通过实例去调用方法
	res2 := res1("abb")
	res3 := res1("abc")
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)

	//函数中的defer：在创建资源时，如数据库连接，文件句柄，锁等，为了在函数执行完毕后及时释放资源，所以有了defer（延时机制）
	res4 := useDefer(222)
	fmt.Println("res4=", res4)
}
