package data2

import (
	"fmt"

	"github.com/Lxb921006/Golang-practise/interface/detail04/data"
)

type Data2 struct{}

func (Data2) ApiPack() {
	fmt.Println("from Data2 ApiPack")
}

func NewData2() *Data2 {
	return &Data2{}
}

func RequireData() {
	data1 := data.NewData()
	data1.DataCheckOne()
}
