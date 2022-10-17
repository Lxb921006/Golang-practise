package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("t=%v, t type=%T\n", t, t)
	fmt.Println("year=", t.Year())
	fmt.Println("month=", int(t.Month()))
	fmt.Println("day=", t.Day())
	fmt.Println("hour=", t.Hour())
	fmt.Println("min=", t.Minute())
	fmt.Println("sec=", t.Second())
	fmt.Println("date=", t.Second())
	//格式化
	fmt.Println("date=", t.Format("2006-01-02 15:04:05")) //必须是这样写，否则出错
	fmt.Println("date=", t.Format("15:04:05"))
	fmt.Println("month=", t.Format("01"))
	// time.Sleep(3 * time.Second) //延时3秒
	time.Sleep(500 * time.Millisecond) //延时0.5秒
	fmt.Println("sleep3")
	fmt.Println("时间戳=", t.Unix())     //Unix获取是的秒数
	fmt.Println("时间戳=", t.UnixNano()) //Unix获取是的纳秒数，
	te := t.Unix()
	tet := time.Unix(te, 0).Format("2006-01-02 15:04:05") //时间戳转时间
	// tef := tet.Format("2006-01-02 15:04:05")
	fmt.Println("tef=", tet)
	test2(test1, 10)

	//builtin函数
	n1 := new(int) //指针类型*int,值默认为0在空间的地址,指针地址，用来分配内存，主要用来分配值类型，比如int,fload32,struct等，返回的是指针
	*n1 = 20
	fmt.Printf("n1 type=%T, n1 val=%v, n1 pointer=%v, pointer val=%v\n", n1, n1, &n1, *n1)
	//make(), 主要是用来分配内存，主要用来分配引用类型，如指针，chan管道，map，slice

	n2 := 100
	var ptr *int = &n2
	fmt.Println("n2=", n2)
	fmt.Println("n2=", &n2)
	fmt.Println("ptr=", ptr)

	*ptr = 200
	fmt.Println("n2-2=", &n2)
	fmt.Println("ptr-2=", ptr)
	fmt.Println("n2-2=", n2)
}

//计算函数的执行时间
type costRun func(n int)

func test2(f costRun, n int) {
	funcName, _, _, _ := runtime.Caller(0)
	start := time.Now().Unix()
	f(n)
	end := time.Now().Unix()
	fmt.Printf("%v run cost time=%d\n", runtime.FuncForPC(funcName).Name(), end-start)
}

func test1(n1 int) {
	n := 0
	for {
		if n < n1 {
			fmt.Printf("n%d\n", n)
		} else {
			break
		}
		n++
	}
}
