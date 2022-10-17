package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//原型模式：用于创建重复的对象，当一个类在创建时开销比较大时，如大量的数据准备，数据库连接等，就可以缓存该对象，当下一次调用，返回该对象的克隆
//原理：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象，通过实现克隆clone()操作，快速生成和原型对象一样的实例

type CPU struct {
	Name string
}

type Rom struct {
	Name string
}

type Disk struct {
	Name string
}

type Computer struct {
	Cpu  CPU
	Rom  Rom
	Disk Disk
}

func (c *Computer) Clone() *Computer {
	resum := *c
	return &resum
}

func (c *Computer) Backup() *Computer {
	pc := new(Computer)
	if err := DeepCopy(pc, c); err != nil {
		panic(err.Error())
	}
	return pc
}

func DeepCopy(dst, src interface{}) (err error) {
	var buf bytes.Buffer
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	cpu := CPU{"i5-1000"}
	rom := Rom{"金士顿16G"}
	disk := Disk{"sam200G"}

	c := Computer{
		Cpu:  cpu,
		Rom:  rom,
		Disk: disk,
	}

	c1 := c.Backup()
	fmt.Println("c1 = ", *c1)
	fmt.Println("c1-1 = ", c1)

}
