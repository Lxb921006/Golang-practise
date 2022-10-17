package main

import "fmt"

type Ainterface interface {
	Say() //自定义的类型调用这个接口，必须实现这个接口的所有方法，否则编译无法通过
}

type Binterface interface {
	Good() //自定义的类型调用这个接口，必须实现这个接口的所有方法，否则编译无法通过，也不能有其他变量存在
	//Name string 会报错
}

type Cinterface interface { //接口也可以继承，但是自定义类型调用这个接口，必须实现继承接口的所有方法，否则报错
	Ainterface //golang低版本接口继承不允许有相同的方法，编译器会报错，高版本可以
	Binterface
}

type Dinterface interface { //空接口，也就是没有任何方法，因此所有数据类型都实现了空接口

}

type Einterface interface { //空接口，也就是没有任何方法，因此所有数据类型都实现了空接口
	Say()
}

type Finterface interface { //空接口，也就是没有任何方法，因此所有数据类型都实现了空接口
	Say()
	Good()
	Hello()
}

type Ginterface interface { //空接口，也就是没有任何方法，因此所有数据类型都实现了空接口
	Einterface
	Finterface
}

type Stu struct {
}

type Hinterface interface { //空接口，也就是没有任何方法，因此所有数据类型都实现了空接口
	Say()
}

type Stu02 struct {
}

func (st *Stu02) Say() {
	fmt.Println("Stu02 Say()")
}

func (st Stu) Say() {
	fmt.Println("Stu Say()")
}

func (st Stu) Good() {
	fmt.Println("Stu Good()")
}

func (st Stu) Hello() {
	fmt.Println("Stu Hello()")
}

type Integer int //不仅仅是结构体，其他数据类型也可以实现接口，并且可以实现多个接口

func (i Integer) Say() {
	fmt.Println("Integer Say()")
}

func (i Integer) Good() {
	fmt.Println("Integer Good()")
}

func main() {
	//var i1 Ainterface 接口不能创建实例，这里会报错，只能通过以下方法来创建，如果Stu结构体没有实现Say()方法也会报错
	var i1 Stu
	var i2 Ainterface = i1
	i2.Say()
	fmt.Println("i2=", i2)

	var i3 Integer = 10
	var i4 Ainterface = i3
	var i5 Binterface = i3
	i4.Say()
	i5.Good()

	i6 := &Stu{}
	var i7 Cinterface = *i6
	i7.Good()
	i7.Say()

	var i8 Dinterface = *i6
	fmt.Println("i8=", i8)

	var i9 interface{} = *i6 //这个写法同上, 空接口可以接受任何数据类型，也就是可以把任何数据类型赋值给空接口，俗称泛型。可以这个非常重要
	fmt.Println("i9=", i9)

	i10 := 12.5
	i9 = i10 //可以把任何数据类型赋值给空接口
	i8 = i10
	i11 := [...]int{1, 2, 3}
	i9 = i11
	fmt.Println("i9=", i9) //i9= 12.5
	fmt.Println("i8=", i8) //i8= 12.5

	var i12 Einterface = i1
	i12.Say()
	var i13 Finterface = i1 //就是只要实现接口的方法就可以（名字必须一样），并不需要关心接口里的方法对应的是谁的
	i13.Good()
	i13.Say()

	var i14 Ginterface = i1
	i14.Hello()
	i14.Good()
	i14.Say()

	//重要
	stu02 := Stu02{}
	var i15 Hinterface = &stu02 //同下写法
	i15.Say()
	// stu02 := &Stu02{}
	// (*stu02).Say()
	// var i15 Hinterface = stu02 //因为Stu02结构体方法参数是指针类型，所以这里必须传Stu02实例的地址
	// i15.Say()
}
