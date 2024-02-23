package main

import (
	"fmt"
)

// 将Student01和Student02共有的字段放到Student里，这个结构体就匿名结构体
type Student struct {
	Name    string
	Age     int
	socre   float64
	Gender  string
	Subject string
}

type Teacher struct {
	Subject string
}

type Student01 struct {
	Student //这里实际省略了一个字段，类型是Student，只是golang底层编译器优化了就不需要了，比如 Stu Student
}

type Student02 struct {
	Student
	Teacher //多重继承，两个匿名结构体如果有相同的字段或者方法，在引用的时候必须明确指定匿名结构体的名字
	Gender  string
}

// 将Student01和Student02共有的方法也绑定到*Student
func (stu *Student) StuInfo() {
	fmt.Printf("name=%v, age=%d, socre=%f\n", stu.Name, stu.Age, stu.socre)
}

func (stu *Student) SetSocre(soc float64) {
	stu.socre = soc
}

func (stu *Student) StudentName() {
	fmt.Println("匿名")
}

func (stu *Student02) StudentName() {
	fmt.Println("stu02")
}

func (s1 *Student01) Testing() {
	fmt.Println("学生1在考试中...")
}

func (s1 *Student02) Testing() {
	fmt.Println("学生2在考试中...")
}

func main() {
	fmt.Println("------------------面向对象编程-继承--------------------")
	fmt.Println("------------------面向对象编程-继承-方式1,直接直接赋值--------------------")
	// stu01 := &Student01{
	// 	Student{
	// 		Name: "lxb",
	// 		Age:  30,
	// 	},
	// }
	// (*stu01).StuInfo()
	// stu01.SetSocre(80.0)
	// stu01.Testing()

	// stu02 := &Student02{
	// 	Student{
	// 		Name:  "lqm",
	// 		Age:   30,
	// 		Socre: 70.3,
	// 	},
	// }
	// (*stu01).StuInfo()
	// stu02.SetSocre(75.6)
	// stu02.Testing()
	fmt.Println("------------------面向对象编程-继承-方式2--------------------")
	stu011 := &Student01{}
	stu011.Student.Name = "lxb" //也可以直接stu011.Name
	stu011.Student.Age = 30
	stu011.Student.socre = 71.3
	stu011.Testing()
	stu011.SetSocre(69.3)
	stu011.StuInfo()
	fmt.Println("stu011 name=", stu011.Name)

	stu022 := &Student02{}
	stu022.Gender = "male" //当结构体跟匿名结构体有相同字段或者方法时，编译器采用就近原则，会优先访问结构体里边的字段或者方法
	stu022.Student.Gender = "famale"
	stu022.Student.Name = "lqm"
	stu022.Student.Age = 30
	stu022.Student.socre = 81.3
	stu022.Teacher.Subject = "语文"
	stu022.Testing()
	stu022.SetSocre(90.3)
	stu022.StuInfo()
	stu022.StudentName()
	stu022.Student.StudentName()

	fmt.Println("stu022 name=", stu022.Name)
	fmt.Println("stu022 Gender=", stu022.Gender)
	fmt.Println("------------------面向对象编程-继承-引入--------------------")
}
