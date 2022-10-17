package main

import (
	"fmt"
	"os"
)

func main() {
	//命令行传入参数,类似python的sys.args,第一个参数是程序名字，第二个开始才是参数名
	p := os.Args
	if len(p) != 2 {
		fmt.Println(len(p))
		panic("参数错误")
	}
	fmt.Println(p)
}
