package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` //新版本优化，序列化的时候字段首字母如果不是大写将会被忽略，而不是像老版本直接报错了
	Age  int
}

func picklePerson() {
	p := Person{
		Name: "lxb",
		Age:  30,
	}

	b, err := json.Marshal(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("picklePerson=", string(b))
}

func unPicklePerson() {
	var p Person

	b, err := json.Marshal(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b=", string(b))

	ub := json.Unmarshal(b, &p)
	if ub != nil {
		fmt.Println(ub)
		return
	}

	fmt.Println("unPicklePerson:", p)
}

func pickleMap() {
	m1 := map[string]string{}
	m1["name"] = "lxb"

	b, err := json.Marshal(&m1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("pickleMap=", string(b))
}

func UnpickleMap() {
	m1 := map[string]string{}
	m1["name"] = "lxb"

	b, err := json.Marshal(&m1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b=", string(b))

	json.Unmarshal(b, &m1)

	fmt.Println("UnpickleMap=", m1)
}

func pickleSlice() {
	s1 := []map[string]interface{}{}
	m1 := map[string]interface{}{}
	m1["name"] = "lxb"
	m1["hobby"] = []string{"篮球", "足球"}
	s1 = append(s1, m1)

	m2 := map[string]interface{}{}
	m2["name"] = "lqm"
	m2["hobby"] = []string{"羽毛球", "足球"}
	s1 = append(s1, m2)

	b, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b=", string(b))

}

func unPickleSlice() {
	s1 := []map[string]interface{}{}
	m1 := map[string]interface{}{}
	m1["name"] = "lxb"
	m1["hobby"] = []string{"篮球", "足球"}
	s1 = append(s1, m1)

	m2 := map[string]interface{}{}
	m2["name"] = "lqm"
	m2["hobby"] = []string{"羽毛球", "足球"}
	s1 = append(s1, m2)

	b, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("b 序列化=", string(b))

	d := []map[string]interface{}{}
	json.Unmarshal(b, &d)

	fmt.Println("unPickleSlice=", d)
}

func main() {
	//反序列化一个json字符串时，要确保反序列化后的数据类型跟原来的序列化前的数据类型一致
	//json序列化
	picklePerson()
	pickleMap()
	pickleSlice()
	fmt.Println("------------------------------")
	//json反序列化
	unPickleSlice()
	unPicklePerson()
	UnpickleMap()

	s := `[{"hobby":["篮球","足球"],"name":"lxb"},{"hobby":["羽毛球","足球"],"name":"lqm"}]`
	s1 := []map[string]interface{}{}

	json.Unmarshal([]byte(s), &s1)

	fmt.Println("s1=", s1)
	fmt.Println("-----------------------------------")
	m1 := map[int]int{}
	m1[1] = 1
	m1[2] = 2
	fmt.Println("m1=", m1)
	for i, v := range m1 {
		fmt.Printf("m[%d]=%d\n", i, v)
	}
}
