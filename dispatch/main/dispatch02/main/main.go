package main

import (
	"fmt"
	"reflect"
)

type Student01 struct {
	Name  string `json:"name"`
	Age   int
	hobby string
}

func (s Student01) StuAge() {
	fmt.Println("-------StuAge-------", s.Age)
}

func (s Student01) StuName(n string, n2 int) (string, int) {
	return s.Name, n2
}

func main() {
	//反射的细节
	//Kind跟type可能时相同也可能是不同,如:传入的是int类型,Kind跟type都是int
	//但是如果传入的是结构体,Kind就是struct,type是包名.结构体名称

	s01 := Student01{
		Name:  "lxb",
		Age:   30,
		hobby: "bas",
	}
	Dispatch03(&s01)
	fmt.Println("s01=", s01) //{lxb 18}
	fmt.Println("----------Dispatch04---------")
	Dispatch04(s01)
	fmt.Println("---------Dispatch05----------")
	Dispatch05(&s01)
	fmt.Println("Dispatch05=", s01)
}

func Dispatch03(v interface{}) {
	v1 := reflect.ValueOf(v)
	//v1.Elem()相当于取到(*s01)
	fmt.Println("v1=", v1.Elem()) //{lxb 30}
	v2 := v1.Elem().Field(1)      // Age的值30
	v2.SetInt(18)                 // 修改Age的值为18
	fmt.Println("v2=", v2)        //18

	v3 := v1.Elem().Method(0)
	if v3.IsValid() {
		fmt.Printf("v3= %v, type= %T\n", v3, v3)
	} else {
		fmt.Println("not exists")
	}
}

func Dispatch04(v interface{}) {
	t1 := reflect.TypeOf(v)
	v1 := reflect.ValueOf(v)

	k1 := v1.Kind()
	if k1 != reflect.Struct {
		fmt.Println("must be a struct type")
		return
	}
	f1 := v1.NumField()
	fmt.Println("struct field number=", f1)
	fmt.Println("field 0=", v1.Field(0))

	//遍历struct字段
	for i := 0; i < f1; i++ {
		//获取tag的值
		tagVal := t1.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field=%d, val=%v, tag=%v\n", i, v1.Field(i), tagVal)
		} else {
			fmt.Printf("field=%d, val=%v\n", i, v1.Field(i))
		}
	}

	m1 := v1.NumMethod()
	//遍历struct方法
	for i := 0; i < m1; i++ {
		fmt.Printf("method=%d, val=%v\n", i, t1.Method(i).Name)
	}

	//调用结构体方法:按首字母排序,在根据index调用结构体方法时,不一定调用的是期望的方法,同理结构体field调用
	// v1.Method(0).Call(nil)

	//给StuName传入参数获取值
	args := []reflect.Value{}
	args = append(args, reflect.ValueOf("lxb"))
	args = append(args, reflect.ValueOf(1))
	res1 := v1.Method(1).Call(args)
	methodName1 := t1.Method(1).Name
	fmt.Printf("%v res1= %v", methodName1, res1[0])

	t1.Method(1).Type.IsVariadic()
}

func Dispatch05(v interface{}) {
	t1 := reflect.TypeOf(v)
	v1 := reflect.ValueOf(v)

	k1 := v1.Kind()
	if k1 != reflect.Ptr || v1.Elem().Kind() != reflect.Struct {
		fmt.Println("want struct type")
		return
	}

	// name := v1.Elem().FieldByName("Name").Interface().(string)
	// fmt.Println("name ===>", name)

	valNum := t1.Elem().NumField()
	for i := 0; i < valNum; i++ {
		if v1.Elem().Field(i).Kind() == reflect.Int {
			v1.Elem().Field(i).SetInt(100)
		}
		tagVal := t1.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d, val= %v, tag=%v\n", i, v1.Elem().Field(i), tagVal)
		} else {
			fmt.Printf("field %d, val= %v\n", i, v1.Elem().Field(i))
		}
	}

	methodNum := t1.Elem().NumMethod()
	for i := 0; i < methodNum; i++ {
		fmt.Printf("method %d, name=%v, arg num = %d\n", i, t1.Elem().Method(i).Name, t1.Elem().Method(i).Type.NumOut())
	}

	// v2 := v1.Interface().(*Student01)

	v1.Elem().FieldByName("Name").SetString("lqm")

	// fmt.Println("v2 = ", v2.Name)

}
