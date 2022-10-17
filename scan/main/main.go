package main

import "fmt"

func main() {
	var name string
	var age byte
	//方式1
	// fmt.Printf("请输入名字:")
	// fmt.Scanln(&name)
	// fmt.Printf("输入名字是:%v", name)
	//方式2
	fmt.Printf("请输入名字,年龄空格隔开")
	fmt.Scanf("%v %d", &name, &age)
	fmt.Printf("输入名字是%v, 年龄是%d", name, age)
}
