package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//类似python的argparse，对传入的参数进行规范解析

	user := flag.String("user", "", "用户名")
	age := flag.Int("age", 0, "年龄")
	file := flag.String("file", "", "文件")
	//重要操作，必须调用该方法，上面的规范参数才生效
	flag.Parse()

	if flag.NFlag() != 3 {
		fmt.Println(flag.ErrHelp)
		return
	}

	os.Remove(*file)

	fmt.Println("hello", *user)
	fmt.Println("age", *age)
}
