package main

import (
	"fmt"
	"reflect"
)

func main() {
	func1 := func(n1, n2 int, name string) {
		fmt.Printf("name = %s func1 val = %d\n", name, n1+n2)
	}

	func2 := func(age int, name string) {
		fmt.Printf("func2 name = %v, age = %d\n", name, age)
	}

	adapter := func(funcName interface{}, args ...interface{}) {
		val := reflect.ValueOf(funcName)
		p := []reflect.Value{}
		for i := 0; i < len(args); i++ {
			p = append(p, reflect.ValueOf(args[i]))
		}
		val.Call(p)
	}

	adapter(func1, 1, 2, "lxb")
	adapter(func2, 30, "lxb")
}
