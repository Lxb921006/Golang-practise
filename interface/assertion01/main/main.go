package main

import "fmt"

//创建一个接口，高聚合，低耦合-多态例子-多态参数，多态数组都可以体现多态
type Usb interface {
	Start() interface{} //接口的方法都没有方法体，也不需要显示的实现
	Stop()
	//Restart() //当调用Usb接口如果这个方法没有被调用，会编译错误，运行报错
}

type Params struct {
	Age int
}

type Phone struct {
	Name string
}

func (p Phone) Start() interface{} {
	fmt.Println("phone start")
	var dd []Params
	d1 := Params{
		Age: 11,
	}
	dd = append(dd, d1)
	return dd
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

func (p Phone) Call() {
	fmt.Println("call phone")
}

type Camera struct {
	Name string
}

func (c Camera) Start() interface{} {
	fmt.Println("camera start")
	var dd []Params
	d1 := Params{
		Age: 10,
	}
	dd = append(dd, d1)
	return dd
}

func (c Camera) Stop() {
	fmt.Println("camera stop")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) { //相当于一个方法实现了不同的功能
	t := usb.Start()
	fmt.Println("t = ", t)
	//使用类型断言解决多态里调用其他类型不存在方法报错的问题
	if phone, info := usb.(Phone); info {
		phone.Call()
	}
	usb.Stop()
}

type Student struct { //Student已经被加载到内存
	Name string
}

func main() {
	//定义Usb数组，方法Phone和Camera结构体
	m1 := [2]Usb{Phone{"苹果"}, Camera{"尼康"}} //必须要实现接口方法方才能把两个结构体放在一个数组里-也就是多态数组
	fmt.Println("m1=", m1)
	computer := Computer{}

	for _, v := range m1 {
		computer.Working(v)
	}

}
