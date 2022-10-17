package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Test01 struct {
	Name string
	Age  int
}

//定义Test01结构体切片类型
type Test01Slice []Test01 //切片里边是放结构体,这样才能放得多

func (t Test01Slice) Len() int {
	return len(t)
}

func (t Test01Slice) Less(i, j int) bool {
	return t[i].Name < t[j].Name
}

func (t Test01Slice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	//接口的实际运用
	s1 := []int{1, 5, 2, 6, 3, 10}
	sort.Ints(s1)
	fmt.Println(s1)

	var s2 Test01Slice
	for i := 0; i < 10; i++ {
		s3 := Test01{
			Name: fmt.Sprintf("lxb%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		s2 = append(s2, s3) //把结构体实例append到切片中
	}
	fmt.Println("----------排序前-----------")
	for _, v := range s2 {
		fmt.Println(v)
	}
	fmt.Println("----------排序后-----------")
	sort.Sort(s2) //因为Test01Slice实现了Sort的三个方法，也就是实现了接口的方法就可以传入该变量，也就是说只要实现接口的所有方法就可以

	for _, v := range s2 {
		fmt.Println(v)
	}

}
