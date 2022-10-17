package pointer //定义包的名字，必须要在第一行

import (
	_ "encoding/json" //可以忽略导入包时，没使用的报错
	"fmt"
	_ "reflect"
	_ "strconv"
	_ "unsafe"
)

//指针类型，形象解析：定义一个i变量会在内存中开辟一个空间同时生成内存地址，用来存储i变量的值如100，然后i变量会指向这个值100，这样在提交程序执行时，程序才能通过内存地址在内存里找到这个i变量对应的值
func main() {
	//基本数据类型在内存布局
	i := 10
	//i的内存地址
	fmt.Println("i的内存地址=", &i)
	//指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	//指针在内存的布局
	//i2是指针变量，类型的是*int，i2的值是&i，只能接受一个内存地址如0xc000006030，且该变量的内存地址类型必须与定义的指针变量类型一直，否则就无法编译通过
	//var a int = 10
	//var prt *flaot64 = &a，这个是错误的
	var i2 *int = &i
	fmt.Println("i2的值=", i2)
	fmt.Println("i2的指针地址=", &i2)
	fmt.Println("i2的指向的值=", *i2)

	//指针案列
	var i3 int = 10
	fmt.Println("i3的地址=", &i3)

	var ptr *int = &i3
	*ptr = 20 //修改这个值时，i3的值也会改变成20，因为ptr自己的内存空间里存的是i3的内存地址，它自己的内存地址指向的是该内存空间，所以修改ptr的值会改变i3原本的值
	fmt.Println("i3的值=", i3)

	var a int = 100
	fmt.Printf("a改变之前的内存地址=%v, a的值=%d\n", &a, a)

	a = 200
	fmt.Printf("a改变之后的内存地址=%v, a的值=%d\n", &a, a)

	var a1 int = 100
	var b1 int = 200
	var ptr1 *int = &a1
	*ptr1 = 300 // a = 300
	ptr1 = &b1  //覆盖了原先的ptr1内存空间存的a1的地址成了的b1的内存地址
	*ptr1 = 500 // a = 300, b = 500, *ptr1 = 500
	fmt.Printf("a1=%d, b1=%d, ptr1=%d", a1, b1, *ptr1)
	//*重要
	//值类型，都有对应的指针类型，形式为*数据类型，如int的对应的指针就是*int以此类推
	//值类型包括：基本的数据类型int系列，bool,string，数组和结构体struct
	//值类型特点：变量直接存储值，内存通常（不是绝对）在栈中分配，如 var i int = 100
	//引用类型：指针，slice切片，map，管道chan，interface等
	//引用类型的特点：变量存储的是一个地址，这个地址对应的空间才是真正的值，内存通常（不是绝对）在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为了一个垃圾由GC回收
	//如指针 var i int = 100, var prt *int = &i

}
