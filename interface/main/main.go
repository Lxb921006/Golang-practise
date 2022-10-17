package main

import "fmt"

//创建一个接口，高聚合，低耦合-多态例子-多态参数，多态数组都可以体现多态
type Usb interface {
	Start() //接口的方法都没有方法体，也不需要显示的实现
	Stop()
	//Restart() //当调用Usb接口如果这个方法没有被调用，会编译错误，运行报错
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("phone start")
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("camera start")
}

func (c Camera) Stop() {
	fmt.Println("camera stop")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) { //相当于一个方法实现了不同的功能
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	//定义Usb数组，方法Phone和Camera结构体
	var m1 [2]Usb //必须要实现接口方法方才能把两个结构体放在一个数组里-也就是多态数组
	m1[0] = phone
	m1[1] = camera
	fmt.Println("m1=", m1)
}
