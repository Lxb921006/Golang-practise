package main

import (
	"fmt"
	"reflect"
)

type Login struct {
	Name string `aaa:"the name =" required:"true"`
	Age  int    `aaa:"the age =" required:"false"`
}

//最终输出如：the name = LXB, the age = 30

func Dispatch(v interface{}) {
	tp := reflect.TypeOf(v).Elem()

	vl := reflect.ValueOf(v).Elem()

	if tp.Kind() != reflect.Struct {
		fmt.Println("类型错误,期望struct")
		return
	}

	fieldNum := vl.NumField()
	for i := 0; i < fieldNum; i++ {
		fval := vl.Field(i)
		fp := fval.Kind()
		ft := tp.Field(i).Tag

		if fp == reflect.String {
			if ft.Get("required") == "true" {
				if fval.String() == "" {
					fmt.Printf("field %v not allow nil\n", tp.Field(i).Name)
					return
				} else {
					fval.SetString("LXB")
					fmt.Printf("%s %s\n", ft.Get("aaa"), fval)
				}
			} else {
				fval.SetString("LXB")
				fmt.Printf("%s %s\n", ft.Get("aaa"), fval)
			}
		} else {
			fmt.Printf("%s %v\n", ft.Get("aaa"), fval.Interface())
		}

	}
}

func main() {
	l := Login{
		Name: "lxb",
		Age:  30,
	}
	Dispatch(&l)
}
