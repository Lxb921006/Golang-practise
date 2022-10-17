package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Student struct {
	Name string
	Age  int
}

type Person struct {
	Student
	Adress string
}

func main() {
	fmt.Println("----------------管道使用------------------")
	//channel管道的使用
	//var c1 chan int
	//c1 = make(chan int, 3)
	c := make(chan int, 3) //这里的容量跟之前的map不一样,map会自动扩容,但是管道不会,设置了多少在cap的输出就是多少，超过管道容量就会报错

	fmt.Println(c)

	//向管道写入数据
	c <- 11
	c <- 22
	c <- 33

	//查看管道长度长度cap(容量)
	fmt.Printf("channel len=%v, cap=%v\n", len(c), cap(c))

	//从管道中读取数据,在没有使用协程情况下,如果管道中的数据取完，也会报错deadlock!
	//先进先出,输出11
	c1 := <-c
	fmt.Println("c1=", c1)
	fmt.Printf("channel len=%v, cap=%v\n", len(c), cap(c))

	c2 := <-c
	fmt.Println("c2=", c2)
	fmt.Printf("channel len=%v, cap=%v\n", len(c), cap(c))

	c3 := <-c
	fmt.Println("c3=", c3)
	fmt.Printf("channel len=%v, cap=%v\n", len(c), cap(c))

	c <- 44
	//去取出数据时也可以不赋值给变量,直接取出
	<-c

	//写入结构体
	cc1 := make(chan Student, 3)
	s1 := Student{
		Name: "lxb",
		Age:  30,
	}
	cc1 <- s1
	s2 := <-cc1
	fmt.Println("s2 Name=", s2.Name)

	ss1 := Student{
		Name: "lqm",
		Age:  30,
	}

	cc1 <- ss1
	ss2 := <-cc1
	fmt.Println("ss2 Name=", ss2.Name)

	//混合类型
	cc2 := make(chan interface{}, 3)
	cc2 <- 10
	cc2 <- "lxb"
	cc2 <- ss1

	cc21 := <-cc2
	cc23 := <-cc2
	cc24 := (<-cc2).(Student)
	fmt.Println("cc21=", cc21)
	fmt.Println("cc23=", cc23)
	//cc24写入的是一个结构体,当取cc24.Name会报错,因为输出cc24类型依然是空接口类型,并没有字段属性,需要用到类型断言
	// if cc3, info := cc24.(Student); info {
	// 	fmt.Println("cc3.Name= ", cc3.Name)
	// } else {
	// 	fmt.Println("不是Student类型")
	// }

	fmt.Println("cc24=", cc24.Name)
	fmt.Println("----------------管道使用细节------------------")
	//在管道取数据时一般都是先进先出
	fmt.Println("----------------管道使用练习------------------")

	pc := make(chan Person, 10)

	for i := 1; i <= cap(pc); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r2 := r.Intn(100)
		p1 := Person{
			Adress:  "广东.河源",
			Student: Student{Name: "lxb" + strconv.Itoa(r2), Age: 30},
		}
		pc <- p1
		p2 := <-pc
		fmt.Println("pc data=", p2)
	}

	fmt.Println("----------------管道的关闭,遍历------------------")
	pc1 := make(chan int, 3)
	pc1 <- 20
	pc1 <- 30
	// close(pc1) //close()是内置,函数专门用来关闭chan关闭后就不能再写入,但是还可以读取
	// close(pc1)
	pp1 := <-pc1 // pp1,ok := <-pc1 可以查看pp1是否取到值，ok是布尔类型
	fmt.Println("pp1=", pp1)

	pc2 := make(chan int, 5)
	pc2 <- 10
	pc2 <- 11

	//这个遍历方式是错的,因为管道是每次取一个就会少一个也就是对应len(pc2)是递减的，如果chan中有两个数据,只会遍历一次,别忘了,管道是引用类型
	// for i := 1; i <= len(pc2); i++ {
	// 	d := <-pc2
	// 	fmt.Println("d=", d)
	// }
	//也不能用容量遍历,容量!=len(chan)

	//正确的遍历方法应该是用for range,并需要关闭管道(在没有使用goroutine的情况下),否则报错：fatal error: all goroutines are asleep - deadlock
	close(pc2)
	for v := range pc2 {
		fmt.Println(v)
	}

}
