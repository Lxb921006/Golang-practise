package main

import (
	"errors"
	"fmt"
)

//golang通过defer，panic，recover来处理异常
func main() {
	test1(10, 0)
	fmt.Println("遇到异常继续")
	//自定义错误
	test3("lxb.conf1")
	fmt.Println("继续执行代码")
}

func test1(n1, n2 int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()
	fmt.Println("res1=", n1/n2)
}

//自定义错误信息
func test2(s string) (err error) {
	if s != "lxb.conf" {
		return errors.New("文件名错误")
	} else {
		return nil
	}
}

func test3(f string) {
	err := test2(f)
	if err != nil {
		//如果错误，就输出错误，并终止程序
		panic(err)
	}
	fmt.Println("继续执行后面代码")
}
