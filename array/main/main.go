package main

import "fmt"

func main() {
	//数组切片:可以存放多个同一类型数据，本身也是一种数据类型，是值类型，在Golang
	a := 1.23232323
	a1 := fmt.Sprintf("%.2f", a)
	fmt.Println("a1=", a1)

	//定义数组
	var a2 [6]float64
	//给数组的每个元素赋值
	total := .0
	for i := 0; i < len(a2); i++ {
		a2[i] = float64(i) + 1.0
		total += a2[i]
		fmt.Printf("a[%d]=%v\n", i, a2[i])
	}
	fmt.Println("a2=", a2)
	fmt.Println("a2-0", a2[0])
	fmt.Println("total=", total)
	//浮点数跟具体整数相除不用转换数据类型，如6它是一个常量，可以当作除数去除于浮点数，但是定义了变量就不行，因为它有数据类型
	fmt.Println("avg1=", total/6)
	fmt.Println("avg2=", total/float64(len(a2)))
	fmt.Println("---------------------------------")
	//数组的第一个元素的地址，就是数组的首地址
	var a3 [3]int //一个int64占用8个字节， int32是4个字节
	fmt.Printf("a3的地址=%p\n", &a3)
	fmt.Println("a3=", a3)
	//a3[1]的地址是前一个地址a3[0]加8
	//数组的各个元素的地址间隔是依据数组的类型决定
	a3[0] = 10
	a3[1] = 20
	a3[2] = 30
	fmt.Println("a3=", a3)
	fmt.Printf("a3[0]的地址=%p\n", &a3[0])
	fmt.Printf("a3[1]的地址=%p\n", &a3[1])
	fmt.Printf("a3[2]的地址=%p\n", &a3[2])
	fmt.Println("------------------初始化数组的方式-------------------")
	var n5 [3]int = [3]int{60, 42, 60}
	var n6 = [3]int{60, 42, 60}
	var n7 = [...]int{86, 103, 105}
	n8 := [...]string{1: "lxb", 2: "lqm"}
	fmt.Println("n5=", n5)
	fmt.Println("n6=", n6)
	fmt.Println("n7=", n7)
	fmt.Println("n8=", n8[0])
	fmt.Printf("n8的类型=%T\n", n8)
	if n8[0] == "" {
		fmt.Println("n8[0]=null")
	} else {
		fmt.Println("n8=", n8[0])
	}
	fmt.Println("------------------数组遍历-------------------")
	n9 := [...]int{10, 11, 12, 13, 14, 15}
	for index, vals := range n9 {
		fmt.Printf("index=%d, vals=%d\n", index, vals)
	}
	fmt.Println("------------------数组注意事项-------------------")
	//数组一旦声明了，长度固定，不能动态变化，元素类型必须一致
	n11 := [3]int{11, 12, 13} //n11[4] = 10 编译不通过
	fmt.Println("n11=", n11)
	//创建数组有默认值，bool数组默认是false，数值数组默认是0，字符串数组默认是"",空
	//var n12 [3]int
	//i := 4
	//fmt.Println("n12=", n12[i]) 报错
	//在golang里，函数传入数组，属于值拷贝，不会影响原来的数组元素，如果确实需要修改传递的数组元素则可以使用指针实现
	//长度是数组类型的一部分，在传递函数参数时，需要考虑数组长度
	n13 := [...]int{11, 22, 33}
	n14 := &n13
	fmt.Println("n14=", *n14)
	n15 := changeArr(n13)
	fmt.Println("n15=", n15)
	fmt.Println("n13=", n13)
	//修改n13[0]的值
	n16 := changeArr02(&n13)
	fmt.Println("n16=", n16)
	fmt.Println("n13-2=", n13)
	//指针数组
	n17, n18 := 55, 66
	var n19 = [2]*int{&n17, &n18}
	fmt.Printf("n19的类型=%T\n", n19)
	fmt.Println("n19=", n19)
	n20 := changeArr03(n19)
	fmt.Println("n20=", *n20[0])
	fmt.Println("-----------------案列----------------")
	var n4 [3]int
	for i := 0; i < len(n4); i++ {
		fmt.Printf("请输入第%d元素:", i+1)
		fmt.Scanln(&n4[i])
	}
	fmt.Println("n4=", n4)

	n10 := returnArr(101)
	fmt.Println("n10=", n10[0])
}

func returnArr(n int) *[1]int { //*[1]int 返回该地址指向的值
	var a = [1]int{n}
	return &a //获取的是地址
}

func changeArr(n [3]int) [3]int {
	n[0] = 12
	return n
}

//通过指针来修改传入的数组元素值
func changeArr02(n *[3]int) [3]int {
	n[0] = 12
	return *n
}

//修改指针数组的元素
func changeArr03(n [2]*int) [2]*int {
	*n[0] = 12
	return n
}
