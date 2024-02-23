package main

import (
	"fmt"
	"reflect"
)

type Movies struct {
	Name  string
	Score int
}

func main() {
	m := Movies{
		Name:  "Golang从入门到放弃",
		Score: 90,
	}
	MoviesPackage()
	fmt.Println("m=", m)
}

func MoviesPackage() {
	//反射还可以创建结构体,需要用到new(),返回一个Value类型的值,该值指向类型为reflect.Type的新申请的零值的指针
	var (
		newMovies *Movies
		newType   reflect.Type
		newElem   reflect.Value
	)
	typ := reflect.TypeOf(newMovies)
	fmt.Println(typ.Kind()) //ptr

	newType = typ.Elem() //typ指向的类型=struct
	fmt.Println(newType.Kind())

	//这里还是返回指针类型,后面操作还得带上Elem()
	newElem = reflect.New(newType)            //返回reflect.Value, 新申请的typ类型的零值指针,指向的地址是*Movies的地址,但是没有数据
	fmt.Printf("newElem type= %T\n", newElem) //reflect.Value

	// newMovies = newElem.Interface().(*Movies)
	// fmt.Printf("newMovies type =%T\n", newMovies)

	// newMovies.Name = "lqm"
	// newMovies.Score = 100
	//newMovies跟newElem都可以修改newElem的值
	newElem.Elem().FieldByName("Name").SetString("lqm")
	newElem.Elem().FieldByName("Score").SetInt(100)

	fmt.Printf("newMovies= %v\n", *newMovies)
}
