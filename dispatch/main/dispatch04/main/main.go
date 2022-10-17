package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%v 完成了减法运算 %d - %d = %d\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func main() {
	c := Cal{}
	CalPackage(&c)
}

func CalPackage(v interface{}) {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	val = val.Elem()
	typ = typ.Elem()

	fieldNum := typ.NumField()
	for i := 0; i < fieldNum; i++ {
		val.Field(i).SetInt(10 - int64(i))
		fmt.Printf("Field %d, Field name = %v, type= %v, val = %d\n", i, typ.Field(i).Name, val.Field(i).Kind(), val.Field(i))
	}

	m1 := []reflect.Value{}
	m1 = append(m1, reflect.ValueOf("lxb"))
	val.Method(0).Call(m1)

}
