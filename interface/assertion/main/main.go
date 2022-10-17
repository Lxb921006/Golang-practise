package main

import "fmt"

type Stu struct {
	Name string
}

type Student struct { //Student已经被加载到内存
	Name string
}

//通过接口和类型断言查看数据类型
func TypeCheck(args ...interface{}) {

	for _, x := range args {
		switch x.(type) { //是关键字，固定写法，类似python的type()
		case bool:
			fmt.Printf("%v是布尔类型\n", x)
		case float64:
			fmt.Printf("%v是float64\n", x)
		case Student:
			fmt.Printf("%v是Student结构体类型\n", x)
		case *Student:
			fmt.Printf("%v是*Student指针类型\n", x)
		default:
			fmt.Printf("%T不知道是什么类型\n", x)
		}

	}
}

func main() {

	var a1 interface{} = Stu{"lxb"}
	var a2 Stu = a1.(Stu) //类型断言-把一个空接口变量赋值给自定义的变量只能用类型推断
	//判断a1是否指向Stu类型的变量，如果是就转成Stu类型，并赋值给a2,否则报错
	//直接赋值给自定义变量会报错如：var a2 Stu = a1
	fmt.Println("a2=", a2)

	var a3 float64 = 10.5
	var a4 interface{} = a3 //空接口可以接收任意的类型的值
	var a5 = a4.(float64)   //类型要匹配，否则报错
	fmt.Printf("a5 的类型=%T\n", a5)

	//检查的类型断言
	a6 := 22.59
	var a7 interface{} = a6
	//转换失败不会报panic
	if a8, info := a7.(float32); info {
		fmt.Printf("a8 的类型=%T\n", a8)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("代码继续走")

	m2 := 10.5
	m4 := true
	m5 := 10
	var m3 interface{} = m2

	//添加对结构体Student，*Student的类型判断
	stu1 := Student{"lqm"}
	//var stu2 *Student = &Student{"lxb"}，以下是省略写法
	stu2 := &Student{"lxb"}

	var m6 interface{} = stu1
	var m7 interface{} = stu2

	TypeCheck(m3, m4, m5, m6, m7)

	fmt.Println("m6=", m6)
	fmt.Println("m7=", m7)
}
