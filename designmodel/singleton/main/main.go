package main

import (
	"fmt"
	"sync"
)

//单例模式：在系统中一个类始终只有一个实例，如mysql连接池，redis连接池等等
type Singleton interface {
	BeginDo()
}

//注意这里是小写
type singleton struct {
	Name string
}

func (s *singleton) BeginDo() {
	fmt.Println("singleton")
}

var (
	once sync.Once
	sin  *singleton
)

//相当于i := 10;var t interface{};t = i
//也就是把singleton的实例赋值给了Singleton接口，因为singleton实现了该接口的方法
func GetSintance() Singleton {
	//只会执行一次
	once.Do(
		func() {
			sin = &singleton{}
		},
	)
	return sin
}

func main() {
	g1 := GetSintance()
	g1.BeginDo()
	fmt.Printf("ptr1 = %p\n", g1)

	g2 := GetSintance()
	fmt.Printf("ptr2 = %p\n", g2)
}
