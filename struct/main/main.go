package main

import (
	"encoding/json"
	"fmt"
)

//字段可以是各种值类型，引用类型
type Cat struct {
	Color string
	Name  string
	Age   int
	Sons  [3]int
}

type Cat02 struct {
	Color, Name string
	Age         int
	Sons        [3]int
	Sister      map[string]string
	Hobby       []string
	ptr         *int
}

type Stu01 struct {
	Name string
	Age  int
}

type Stu03 struct {
	Name string
	Age  int
}

type Stu04 struct {
	Name string
	Age  int
}

type Test01 struct {
	x, y int
}

type Test02 struct {
	left, right Test01
}

type Test03 struct {
	left, right *Test01
}

type Test04 struct {
	Num  int
	Name string
}

//struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景：序列化和反序列化
type Test05 struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
}

//相当于Test05结构体的别名，但是不等同于Test05，在golang里Test06就是一个新的数据类型
type Test06 Test05

func main() {
	fmt.Println("--------------------struct定义----------------------")
	//golang面向“对象”编程，实际上在golang里边并没有类的概念，通过结构体来替代
	struct01("cat02")
	var cat01 Cat
	cat01.Name = "lxb"
	cat01.Color = "红色"
	cat01.Age = 4
	cat01.Sons = [...]int{10, 11, 12}
	fmt.Println("cat01=", cat01)

	fmt.Println("--------------------struct使用事项----------------------")
	var cat02 Cat02
	if cat02.Hobby == nil {
		fmt.Println("cat02.Hobby=none")
	}
	fmt.Println("cat02=", cat02)
	//结构体的字段类型是：指针，slice，map，默认值都是nil，也就是还没有分配空间,需要make来分配空间后才能使用
	// cat02.Hobby[0] = "lxb"  报错，必须要先make
	cat02.Hobby = make([]string, 10)
	cat02.Hobby[0] = "football"
	s1 := 10
	cat02.ptr = &s1
	fmt.Println("cat02=", cat02)
	//不通结构体的变量的字段是独立的，互不影响，一个结构体变变量字段的更改，不会影响另一个结构体变量，因为结构体是值类型
	var stu01 Stu01
	stu01.Name = "lxb"
	stu02 := stu01
	stu02.Name = "lqm"
	fmt.Println("stu01=", stu01)
	fmt.Println("stu02=", stu02)
	fmt.Println("stu01=", stu01)
	fmt.Println("--------------------struct实例的创建----------------------")
	var stu03 Stu03 = Stu03{} //开发中推荐这个写法
	stu04 := Stu03{}
	stu03.Age = 100
	stu04.Name = "lxb"
	stu05 := Stu03{"lqm", 30}
	var stu06 *Stu03 = new(Stu03)
	stu06.Name = "lxb92" //指针类型也可以这样访问的原因是golang底层已经做了转换即：stu06.Name = "lxb92" 转换成(*stu06).Name =  "lxb92"
	stu07 := new(Stu03)  //等同上，返回的是一个指针
	stu07.Name = "lxb9222"
	stu08 := &Stu03{}
	stu09 := &Stu03{"lxb222999", 30}
	stu08.Name = "lxb92222222"
	fmt.Println("stu04=", stu04)
	fmt.Println("stu03=", stu03)
	fmt.Println("stu05=", stu05)
	fmt.Println("stu06=", *stu06)
	fmt.Println("stu07=", *stu07)
	fmt.Println("stu08=", *stu08)
	fmt.Println("stu09=", *stu09)
	stu10 := Stu04{}
	stu10.Age = 10
	stu10.Name = "lxb"
	fmt.Printf("stu10-1=%v\n", stu10)
	stu11 := &stu10
	fmt.Printf("stu10-1=%v\n", stu10)
	(*stu11).Name = "lqm"
	fmt.Printf("stu10的地址=%p\n", &stu10)
	fmt.Printf("stu11的地址=%p, 值等于=%p\n", &stu11, stu11)
	fmt.Printf("stu10-2=%v\n", stu10)
	fmt.Println("--------------------struct使用细节----------------------")
	//结构体的所有字段在内存中是连续分布的，可以通过地址的加减来获取下一个字段的值，速度快
	s2 := Test02{Test01{1, 2}, Test01{3, 4}}
	fmt.Println("s2=", s2)
	fmt.Printf("s2.left.x的地址=%p\n", &s2.left.x)
	fmt.Printf("s2.left.y的地址=%p\n", &s2.left.y)
	fmt.Printf("s2.left的地址=%p\n", &s2.left)
	fmt.Printf("s2.right.x的地址=%p\n", &s2.right.x)
	fmt.Printf("s2.right.y的地址=%p\n", &s2.right.y)
	fmt.Printf("s2.right的地址=%p\n", &s2.right)
	//指针类型，字段值本身地址是连续的，但是他们指向的地址不一定是连续的
	s3 := Test03{&Test01{10, 20}, &Test01{30, 40}}
	fmt.Printf("s3.left本身的地址=%p\n", &s3.left)
	fmt.Printf("s3.right本身的地址=%p\n", &s3.right)
	fmt.Printf("s3.left指向的地址=%p\n", s3.left)
	fmt.Printf("s3.right指向的地址=%p\n", s3.right)
	//结构体和其他类型进行转换时，需要有完全相同的字段（字段(字段的名字)，字段的顺序，个数和类型）
	// s4 := Test04{}
	// s5 := Test05{}
	// s4 = Test04(s5) //强转
	// fmt.Println("s4=", s4)
	// fmt.Println("s5=", s5)
	// //结构体进行type重新定义，相当于别名，Golang认为是新的数据累心通过，但是互相可以强转
	// s6 := Test06{}
	//s6 = s5 这里会报错，即使是别名也不能跟原来的结构体相等，golang认为是新的数据类型，只能通过强转，如上下所示
	//结论就是不管type自定义什么类型的别名都不等于原来的数据类型，都只能通过强转来实现 如，
	//type newint int, var i new int=10;var t int = 10;i=t?这个是错误的，只能是i=newint(t)强转实现
	//其他同上
	// s6 = Test06(s5)
	// fmt.Println("s6=", s6)

	//struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景：序列化和反序列化，服务端往客户端传值，一般都是传变量的小写，如果传大写会增加沟通成本
	//将s7序列化成json格式字符串, 结构体json序列化，字段首字母必须是大写，否则json后只能输出空串，因为golang里边想要在别的包引用结构体字段，字段首字母就必须是大写才可以
	//json.Marshal原理用到的是反射
	s7 := Test05{1, "lxb"}
	jsonS7, err := json.Marshal(s7)
	if err != nil {
		fmt.Println("json error=", err)
	}
	fmt.Println("jsonS7=", string(jsonS7))
	fmt.Printf("jsonS7=%q\n", jsonS7)
	fmt.Printf("jsonS7 type=%T\n", jsonS7)
}

func struct01(s string) { //没有使用结构体，写法非常不方便，鸡肋
	cat := make(map[string]map[string]string, 5)
	cat["cat01"] = make(map[string]string)
	cat["cat01"]["color"] = "red"
	cat["cat01"]["age"] = "40"
	cat["cat01"]["name"] = "cat01"
	if cat[s] != nil {
		fmt.Println("名字=", s)
		for k, v := range cat[s] {
			fmt.Printf("%v=%v\n", k, v)
		}
	} else {
		fmt.Printf("这只猫%v不存在\n", s)
	}
}
