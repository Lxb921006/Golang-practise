package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Hobby string
}

type Monster struct {
	Name  string
	Hobby string
}

func (s Student) Skill() {
	fmt.Println(s.Hobby)
}

func main() {
	//反射入门:需要用到reflect,reflect.TypeOf,reflect.ValueOf
	s1 := Student{
		Name:  "lxb",
		Hobby: "basketball",
	}
	m1 := Monster{
		Name:  "猴哥",
		Hobby: "爬树",
	}

	Dispatch01(s1)
	fmt.Println("-------------------")
	Dispatch01(&m1)
}

func Dispatch01(v interface{}) {
	t1 := reflect.TypeOf(v)
	fmt.Println("t1= ", t1)
	//main.Student

	k1 := t1.Kind() //struct

	fmt.Printf("v5= %v, type= %T\n", k1, k1)

	n1 := t1.Name()
	fmt.Println("n1= ", n1)

	v1 := reflect.ValueOf(v)
	fmt.Printf("v1= %v, type= %T\n", v1, v1) //{lxb basketball}

	k2 := v1.Kind()
	fmt.Printf("k2= %v, type= %T\n", k2, k2)

	v2 := v1.Field(0)
	fmt.Printf("v2= %v, type= %T\n", v2, v2) //lxb 虽然输出跟v3的结果一样,但是v2,v3两个是不同的类型v2= lxb, v2 type= reflect.Value

	v3 := v2.String()
	fmt.Printf("v3= %v, type= %T\n", v3, v3) //lxb v3= lxb, v3 type= string

	//在转回原本的类型之前需要转成接口类型,虽然类型是Student,但是在编译之前调用Stu结构体字段信息会报错,是运行时的反射
	v4 := v1.Interface()
	fmt.Printf("v4= %v, type=%T\n", v4, v4) // 所以这里v.Hobby编译器无法通过

	//转回原本的类型需要用到类型断言

	switch v5 := v4.(type) {
	case Student:
		fmt.Println(v5.Name)
	case Monster:
		fmt.Println(v5.Name)
	default:
		fmt.Println("未知类型")
	}

}
