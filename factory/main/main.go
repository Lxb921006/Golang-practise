package main

import (
	"fmt"
	model "lxb-learn/factory/model"
) //给包起个别名

func main() {
	fmt.Println("------------------工厂模式--------------------")
	//场景：当定义个结构体变量名首字母是小写时，但是又要可以在别的包里边可以创建这个结构体的实例，这里就需要用到工厂模式
	// s := model.Student{
	// 	Name: "lxb",
	// 	Age:  30,
	// }
	s := model.Newstudent("lxb", 30)
	fmt.Println("s=", *s)
	fmt.Println("s=", s.Newage()) // 等同于(*s).Newage()

	fmt.Println("-----------------------------")

}
