package main

import "fmt"

//建造者模式：使用多个简单的对象一步一步构建成一个复杂的对象
type Builder interface {
	buildDisk()
	buildCPU()
	buildRom()
}

type SuperComputer struct {
	Name string
}

func (*SuperComputer) buildDisk() {
	fmt.Println("硬盘")
}

func (*SuperComputer) buildCPU() {
	fmt.Println("cpu")
}

func (*SuperComputer) buildRom() {
	fmt.Println("内存")
}

type LowComputer struct {
	Name string
}

func (*LowComputer) buildDisk() {
	fmt.Println("硬盘2")
}

func (*LowComputer) buildCPU() {
	fmt.Println("cpu2")
}

func (*LowComputer) buildRom() {
	fmt.Println("内存2")
}

type Drictor struct {
	builder Builder
}

func (d *Drictor) Consturck() {
	d.builder.buildCPU()
	d.builder.buildDisk()
	d.builder.buildRom()
}

func NewDrictor(b Builder) *Drictor {
	return &Drictor{
		builder: b,
	}
}

func main() {
	sc := SuperComputer{}
	lc := LowComputer{}
	d := NewDrictor(&sc)
	d.Consturck()
	d2 := NewDrictor(&lc)
	d2.Consturck()
}
