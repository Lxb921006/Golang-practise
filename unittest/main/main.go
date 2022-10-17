package main

import "fmt"

func main() {
	//单元测试：在工作中，需要确认一个函数或者模块是否正确
	//传统测试方法方法
	//查看编写的add函数是否正确
	//这里不需要main也能执行的原因是：testing框架会隐藏main(),然后在import xxx_test.go,再放到main()运行
	//单元测试命名必须是Test开头

	//如果只想测试单个文件，一定要带上被测试的文件名，默认情况下，testing框架会扫描整个目录
	//如：go test -v add_test.go unittest02.go
	//测试单个方法：go test -v -test.run TestAdd
	//测试整个包：go test -v
	res := add(5)
	fmt.Println("res=", res)

	fmt.Println("------------------单元测试------------------")

}

func add(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
