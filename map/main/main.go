package main

import (
	"fmt"
	"sort"
)

//提前体验struct结构体
type useStruct struct {
	name string
	age  int
}

func main() {
	//map是一种key-value的数据结构，也称为关联字段或数组，映射，类似其他编程语言的集合
	//map声明是不会分配内存的，初始化需要make，分配内存后才能赋值，使用，作用是给map分配内存空间
	//key不能重复，会覆盖原来的val
	//golang里，map是无序的
	fmt.Println("--------------------map(映射)定义方式----------------------")
	//var m1 map[int]string //这样定一个map，需要make初始化
	// var m2 map[string]string
	m1 := make(map[int]string, 2)          //分配可以放2对的key-val
	m2 := make(map[string]string)          //可以直接用make初始化map然后赋值给变量
	m3 := map[string]string{"name": "lxb"} //声明的时候直接赋值
	m33 := map[string]string{}
	m33["name"] = "lxb"
	m1[222] = "lqm"
	m1[111] = "lxb"
	m1[222] = "lqmiii"
	m1[333] = "lyy"
	m1[333] = "lyy2222" //会覆盖上面的lyy->lyy2222
	m1[111] = "lxb888"
	m2["man02"] = "lxb"
	m2["man01"] = "lqm"
	m3["age"] = "30"
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)
	fmt.Println("m3=", m3)
	fmt.Println("m33=", m33)
	fmt.Println("--------------------map(映射)案列----------------------")
	m4 := make(map[string]map[string]string)
	m5 := make(map[string]map[string][2]string)
	m4["stu01"] = make(map[string]string)
	m4["stu01"]["gender"] = "男"
	m4["stu01"]["name"] = "lxb"
	// m4["stu01"] = map[string]string{"name": "lxb", "age": "30"}
	// m4["stu02"] = map[string]string{"name": "lqm", "age": "30"}
	m5["arr01"] = map[string][2]string{"age": {"30", "31"}}
	m5["arr01"]["name"] = [2]string{"lxb", "lqm"}
	m4["stu01"]["name"] = "lxb92"
	//map删除
	delete(m4, "stu02")
	//如果要删除全部key，则为m4重新make一个空间即可，如：m4 = make(map[string]map[string]string)
	//m4 = make(map[string]map[string]string)
	//map的key查找
	val, res := m4["stu02"]
	if res {
		fmt.Println("val=", val)
	} else {
		fmt.Println("val=none")
	}
	fmt.Println("m4=", m4["stu01"])
	fmt.Println("m5=", m5)
	fmt.Println("--------------------map遍历----------------------")
	m6 := map[string]string{"name": "lxb", "age": "30"}
	for k, v := range m6 {
		fmt.Printf("key=%v, val=%v\n", k, v)
	}
	m7 := make(map[string]map[string]string)
	m7["name1"] = map[string]string{"name1": "lxb1", "age1": "30"}
	m7["name2"] = map[string]string{"name2": "lxb2", "age2": "30"}
	fmt.Println("m7=", m7)
	for k1, v1 := range m7 {
		fmt.Printf("key1=%v, val1=%v\n", k1, v1)
		for k2, v2 := range v1 {
			fmt.Printf("key2=%v, val2=%v\n", k2, v2)
		}
	}
	fmt.Println("--------------------map长度----------------------")
	m8 := make(map[string]string)
	m8["name"] = "lxb"
	m8["age"] = "30"
	m8["id"] = "1"
	fmt.Println("m8长度=", len(m8))
	fmt.Println("--------------------map切片----------------------")
	m9 := make([]map[string]string, 2)
	m9[0] = map[string]string{"name": "lxb"}
	m9[1] = map[string]string{"name2": "lxb2"}
	fmt.Println("m9-1=", m9)
	//动态添加map append()
	m10 := map[string]string{"name3": "lxb3"}
	m9 = append(m9, m10)
	fmt.Println("m9-2=", m9)
	fmt.Println("--------------------map排序----------------------")
	m11 := map[int]int{1: 1, 3: 3, 2: 2}
	m12 := []int{}
	fmt.Println("m12=", m12)
	fmt.Println("m11=", m11)
	//遍历输出还是无序
	for k := range m11 {
		m12 = append(m12, k)
	}
	fmt.Println("m12=", m12)
	sort.Ints(m12) //小到大排序
	fmt.Println("m12=", m12)
	for i := 0; i < len(m12); i++ {
		fmt.Printf("m11[%d]=%d\n", m12[i], m11[m12[i]])
	}
	fmt.Println("--------------------map注意事项----------------------")
	m13 := make(map[int]string, 2)
	m13[1] = "lxb"
	m13[2] = "lxb2"
	m13[3] = "lxb3"
	fmt.Println("m13-1=", m13)
	mapChange01(m13)
	fmt.Println("m13-2=", m13)

	m15 := make(map[string]useStruct)
	stu01 := useStruct{
		"lxb",
		30,
	}
	m15["stu01"] = stu01
	fmt.Println("m15=", m15["stu01"].name)

	for k, v := range m15 {
		fmt.Printf("k=%v, v=%v\n", k, v.age)
	}
}

func mapChange01(m map[int]string) {
	m[1] = "lqm"
}
