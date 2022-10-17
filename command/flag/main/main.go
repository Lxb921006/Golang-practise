package main

import (
	"flag"
	"fmt"
)

func main() {
	//类似python的argparse，对传入的参数进行规范解析
	user := ""
	age := 0

	flag.StringVar(&user, "user", "", "用户名")
	flag.IntVar(&age, "age", 0, "年龄")

	//重要操作，必须调用该方法，上面的规范参数才生效
	flag.Parse()

	fmt.Println("hello", user)
	fmt.Println("age", age)
}
