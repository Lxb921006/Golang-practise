package main

import "fmt"

type Stu01 struct {
	Name string
	Age  int
}

func (m *Stu02) StuMethod() {
	fmt.Println("Stu02")
}

type Stu02 struct {
	Name string
	Age  int
}

func (m *Stu01) StuMethod() {
	fmt.Println("Stu01")
}

type Stu05 struct {
	Name string
}

type Stu struct { //多重继承，两个匿名结构体如果有相同的字段或者方法，在引用的时候必须明确指定匿名结构体的名字
	*Stu01 //也可以是指针类型
	*Stu02
}

type Stu2 struct { //多重继承，两个匿名结构体如果有相同的字段或者方法，在引用的时候必须明确指定匿名结构体的名字, 不建议多重继承
	Stu01
	Stu02
}

type Stu03 struct {
	s Stu01 //嵌套有名结构体，也就是组合关系
}

type Stu04 struct {
	Stu02
	int //基本数据类型也可以当作匿名字段
}

type Stu3 struct { //多重继承，两个匿名结构体如果有相同的字段或者方法，在引用的时候必须明确指定匿名结构体的名字, 不建议多重继承
	Stu01
	Stu05
}

func main() {
	//匿名结构体是指针类型演示
	st := Stu{
		&Stu01{
			Name: "lxb",
			Age:  30,
		},
		&Stu02{
			Name: "lqm",
			Age:  30,
		},
	}
	fmt.Println("st1=", *st.Stu01)

	st2 := &Stu2{}
	//st2.Age = "30" //错误写法，编译无法通过，除非在Stu这个结构体里边也存在Age属性，对方法也是一样
	(*st2).Stu01.Age = 30
	st.Stu01.StuMethod() //嵌套多个匿名结构体,如果没有

	//有名结构体演示
	st3 := &Stu03{}
	st3.s.Age = 30 //有名结构体必须带上Stu03结构体的字段名字才能引用Stu01结构体的Age属性，除非Stu03里边有Age属性
	fmt.Println("st3=", *st3)

	//匿名字段是基本数据类型的使用方式
	st4 := &Stu04{}
	st4.Name = "lxb-3"
	st4.Age = 600
	st4.int = 300
	fmt.Println("st4=", *st4)

	st5 := &Stu3{}
	st5.Age = 50
	fmt.Println("st5=", *st5)
}
