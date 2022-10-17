package main //定义包的名字，必须要在第一行，package的名字尽量跟当前目录的名字一致

import (
	"encoding/binary"
	_ "encoding/json" //可以忽略导入包时，没使用的报错
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

//什么是栈？
//主要表现:递归(先入后出)
//比如说一个慈善家建了一栋楼(栈),没地方睡的人(局部变量)都可以去那里度过一晚上,睡一晚了就走了.
//但是在旅馆(堆), 你不能随便就进去睡一晚,要先给钱拿到钥匙(new申请空间),然后睡完了还要把钥匙还回去(释放空间);
//栈， 是硬件， 主要作用表现为一种数据结构， 是只能在一端插入和删除数据的特殊线性表。允许进行插入和删除操作的一端称为栈顶， 另一端为栈底。栈按照后进先出的原则存储数据， 最先进入的数据被压入栈底， 最后进入的数据在栈顶， 需要读数据时从栈顶开始弹出数据。栈底固定， 而栈顶浮动。栈中元素个数为零时称为空栈。插入一般称为进栈(push) , 删除则称为出栈(pop) 。栈也被称为先进后出表， 在函数调用的时候用于存储断点， 在递归时也要用到栈。
//
//在计算机系统中， 栈则是一个具有以上属性的动态内存区域。程序可以将数据压入栈中， 也可以将数据从栈顶弹出。在i386机器中， 栈顶由称为esp的寄存器进行定位。压栈的操作使栈顶的地址减小， 弹出的操作使栈顶的地址增大。
//
//栈在程序的运行中有着举足轻重的作用。最重要的是， 栈保存了一个函数调用时所需要的维护信息， 这常常被称为堆栈帧。栈一般包含以下两方面的信息：
//1） 函数的返回地址和参数。
//2） 临时变量：包括函数的非静态局部变量及编译器自动生成的其他临时变量。
//
//堆， 是一种动态存储结构， 实际上就是数据段中的自由存储区， 它是C语言中使用的一种名称， 常常用于存储、 分配动态数据。堆中存入的数据地址向增加方向变动。堆可以不断进行分配直到没有堆空间为止， 也可以随时进行释放、 再分配， 不存在顺序问题。
//
//堆内存的分配常通过malloc() 、 calloc() 、 realloc() 三个函数来实现。而堆内存的释放则使用free() 函数。
//堆和栈在使用时“生长”方向相反， 栈向低地址方向“生长”， 而堆向高地址方向“生长”。

//https://studygolang.com/pkgdoc Golang的官方包使用解析

//学习打卡：整数类型P37

//int8:   -128 ~ 127
// int16:  -32768 ~ 32767
// int32:  -2147483648 ~ 2147483647
// int64:  -9223372036854775808 ~ 9223372036854775807

//无符号
// uint8:  0 ~ 255
// uint16: 0 ~ 65535
// uint32: 0 ~ 4294967295
// uint64: 0 ~ 18446744073709551615

//声明全局变量，并初始化，%T打印类型，%v原始输出，%d数字，%c打印码值，%f浮点数
//Golang标识符有严格的大小写区分

//int8: 8位,就是一个字节,1bytes=8bit
//int16: 2个字节
//int32: 4个字节
//int64:8个字节
//
//float32: 4个字节
//float64: 8 个字节
//
//int 比较特殊,占用多大取决于你的cpu
//32位cpu 就是 4个字节
//64位 就是 8 个字节
//
//float32: 4个字节
//float64:8个字节

//1B（byte，字节）= 8 bit（位）

type User struct {
	Name string
	Age  int
}

type UserMap map[int]User

const (
	c103 = iota //这里表示给c103赋值0,后面c104在c103的基础上+1以此类推 0
	c104        // 1
	c105 = iota //0
	c106        //1
)

func main() {
	fmt.Println("adasdad")

	//标准声明方式
	var name string
	var age int
	var gender bool
	fmt.Println("var---->", name, age, gender)

	//批量声明
	var (
		a string
		b int
		c bool
	)
	fmt.Println("var1---->", a, b, c)

	//声明变量同时指定初始值
	var name1 string = "lxb"
	var age1 int = 30
	var gender1 bool = true
	fmt.Println("var1.1---->", name1, age1, gender1)

	//短变量声明，只能用在函数内部，后面会常用
	m, n := 10, "lxb"
	fmt.Println("var3---->", m, n)
	fmt.Println("var3---->", m, n)
	//test
	name4 := "lxb"
	fmt.Println("name4 type-->", reflect.TypeOf(name4))

	//test2
	var mm = 100
	fmt.Println("mm-->", mm)
	mm = 202
	fmt.Println("mm-->", mm)
	m1, m2 := "100", "200"
	fmt.Println(m1 + m2)

	//test3
	var jj uint8 = 255
	fmt.Println("jj-->", jj)
	//test4
	var kk int8 = 127
	fmt.Println("kk--->", kk)

	//test5数据类型查看
	var name5 = "廖旋彪"
	fmt.Println("name5--->", name5)
	fmt.Println("name5 type--->", reflect.TypeOf(name5))
	fmt.Printf("name5 type--->%T, name5占用的字节数--->%d\n", name5, unsafe.Sizeof(name5))

	//test fload
	var ff = 0.2500001
	fmt.Println("ff--->", ff)
	fmt.Printf("ff type--->%T\n", ff)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", ff), 64)
	fmt.Print(value)

	//test 科学计数法
	num := 5.1234e2   //相当于5.1234*10的2次方，也就是5.1234*100=512.34，大写E也是同等效果
	num2 := 5.1234e-2 //相当于5.1234/(除于)10的2次方，也就是5.1234/100=0.051234，大写E也是同等效果
	fmt.Println("num--->", num, num2)

	//test 十进制
	nn := .512
	fmt.Println("十进制-->", nn)

	//字符类型：Golang中没有专门的字符类型，如果要存储单个字符（字母），一般用byte
	//Golang里双引号表示字符串，单引号表示字符，byte类型默认输出的是对应字符ASCLL表里的码值
	var cc1 byte = 'a'
	var cc2 byte = '1'
	fmt.Println("cc1--->", cc1, ",cc2--->", cc2)

	//输出对应的值，需要格式化输出
	fmt.Printf("cc1--->%c, cc2--->%c\n", cc1, cc2)
	var cc3 int = '廖'
	//var cc3 byte = '廖' 这个会报错溢出，因为超出了byte的范围，可以改成如上输出
	fmt.Printf("cc3--->%c, cc3的码值--->%d\n", cc3, cc3)
	cc4, cc6 := 2, "3"
	cc5 := "%." + strconv.Itoa(cc4) + "f" // strconv.Itoa(11), int转字符串，strconv.Atoi("777")字符串转int
	cc7, _ := strconv.Atoi(cc6)
	fmt.Println("cc5--->", cc5, "cc7--->", reflect.TypeOf(cc7))

	//在Go中，字符本质是一个整数，直接输出时时该字符对应的utf-8码值
	fmt.Printf("cc3--->%c, cc3的码值--->%d\n", cc3, cc3)

	//当个变量赋值一个数字，然后格式化输出%c时，会输出该数字对应的Unicode字符
	cc8 := 97
	fmt.Printf("cc8--->%c\n", cc8)

	//字符是可以运算的，相当于一个整数，运输时是按照码值运行
	cc9 := 10 + 'a'
	fmt.Println("cc9--->", cc9)

	//Golang中允许使用转义字符'\'来将其后的字符转变为特殊字符型常量
	cc10 := '\n'
	fmt.Println("cc10--->", cc10)

	//布尔值，默认是false
	vv1 := false
	vv2 := true
	fmt.Println("vv1--->", vv1, "vv2--->", vv2)

	//字符串类型：Golang是由字节组成,一旦定义不能修改
	str1 := "abcd"
	intt1 := 100
	intt1 = 101
	fmt.Println("str1--->", str1, "intt1--->", intt1)

	//反引号,使字符串原生形式输出包括特殊字符，可以防止攻击，输出源代码
	str2 := `const n5 = iota //0

	//定义数量级
	const ()
	
	func main() {
		fmt.Println(n1, n2, n3, n4, n5)
	}`
	fmt.Println("str2--->", str2)
	//字符串拼接
	str3 := "aaa" +
		"bbb"
	str3 += "ccc"
	fmt.Println("str3--->", str3)

	//基本数据类型的转换，不同类型的类型变量之间赋值时需要显式转换，不会自动转换
	yy := 100
	var yy1 float64 = float64(yy) //将yy的int类型转换成float64类型赋值给yy1,必须显式的转换，var yy1 float64 = yy这个是错误的
	fmt.Println("yy1-->", reflect.TypeOf(yy1))

	//在转换找那个，比如将int64转成int8(-128-127)，编译时不会报错，只是转换的结果按溢出处理，和我们希望的结果不一样,所以不能随便乱转换，需要考虑范围问题
	var yy3 int64 = 8888
	var yy4 int8 = int8(yy3)
	fmt.Println("yy4-->", yy4)

	//转换练习
	// var tt1 int32 = 12
	// var tt2 int64
	// var tt3 int8

	// tt2 = tt1 + 20 //tt1是int32类型，加上20，赋值给类型为int64的tt2是错误的，同下
	// //tt2 = int64(tt1) + 20 这个才是正确
	// tt3 = tt1 + 20
	// //tt3 = int8(tt1) + 20
	// fmt.Println("tt2-->%v, tt3-->%v", tt2, tt3)

	// var tt1 int32 = 12
	// var tt2 int8
	// var tt3 int8

	// tt2 = int8(tt1) + 127 //这里可以编译通过，但是当做会溢出处理，结果不一致
	// tt3 = int8(tt1) + 128 //这里+128直接超出int8的范围，编译不会通过
	// fmt.Println("tt2-->%v, tt3-->%v", tt2, tt3)

	//基本数据类型和string的转换
	rr1 := 100
	rr2 := 1.2345
	var rr3 byte = 'l'
	rr4 := true
	var strr1 string
	var strr2 string

	//方式1：fmt.Sprintf
	strr1 = fmt.Sprintf("%d", rr1)
	fmt.Printf("strr1 type-->%T, strr1 res -->%q\n", strr1, strr1)

	strr1 = fmt.Sprintf("%f", rr2)
	fmt.Printf("strr1 type-->%T, strr1 res -->%q\n", strr1, strr1)

	strr1 = fmt.Sprintf("%t", rr4)
	fmt.Printf("strr1 type-->%T, strr1 res -->%q\n", strr1, strr1)

	strr1 = fmt.Sprintf("%c", rr3)
	fmt.Printf("strr1 type-->%T, strr1 res -->%q\n", strr1, strr1)

	//方式2: strconv
	strr2 = strconv.FormatInt(int64(rr1), 10)
	fmt.Printf("strr2 type-->%T, strr2 res -->%q\n", strr2, strr2)

	strr2 = strconv.FormatFloat(rr2, 'f', 10, 64)
	fmt.Printf("strr2 type-->%T, strr2 res -->%q\n", strr2, strr2)

	strr2 = strconv.FormatBool(rr4)
	fmt.Printf("strr2 type-->%T, strr2 res -->%q\n", strr2, strr2)

	strr2 = strconv.FormatUint(uint64(rr3), 10)
	fmt.Printf("strr2 type-->%T, strr2 res -->%q\n", strr2, strr2)

	//string转基本数据类型
	rr5 := "true"
	var rr6 bool
	rr6, _ = strconv.ParseBool(rr5)
	fmt.Printf("rr6 type-->%T, rr6 res -->%t\n", rr6, rr6)

	var rr7 string = "1000000"
	var rr8 int64
	rr8, _ = strconv.ParseInt(rr7, 10, 64) //非要转成int类型，直接显式转换就好，如下一行代码所示
	var rr9 int = int(rr8)
	fmt.Printf("rr8 type-->%T, rr8 res -->%d\n", rr8, rr8)
	fmt.Printf("rr9 type-->%T, rr9 res -->%d\n", rr9, rr9)

	rr10 := "1.123456"
	var rr11 float64
	rr11, _ = strconv.ParseFloat(rr10, 64)
	fmt.Printf("rr11 type-->%T, rr11 res -->%f\n", rr11, rr11)

	// var rr12 string = "a"
	// var rr13 uint64
	// rr13, _ = strconv.ParseUint(rr12, 10, 64)
	// fmt.Printf("rr13 type-->%T, rr13 res -->%d\n", rr13, rr13)

	var rr14 int = '中'
	fmt.Printf("rr14-->%T, rr14 res -->%c, rr14的ascll码值-->%d\n", rr14, rr14, rr14)

	rr15 := "lxb" //字符串转bool值，如果不是true或者false，结果都是false，即使rr16 = true，结果也是一样，其他类型也是类似这种道理
	var rr16 bool
	rr16, _ = strconv.ParseBool(rr15)
	fmt.Printf("rr16 type-->%T, rr16 res -->%t\n", rr16, rr16)
	ErrorInfo := "err"
	fmt.Println("ErrorInfo=", ErrorInfo)

	//数组
	vai := []int{0, 2, 3}
	vas := [...]string{"lxb", "1"}
	for i := 0; i < len(vai); i++ {
		fmt.Println(vai[i])
	}
	fmt.Println("len vas=", len(vas))
	for i := 0; i < len(vas); i++ {
		fmt.Println(vas[i])
	}
	ss := make([]string, 5)
	fmt.Println("ss=", ss)

	fmt.Println("-------------常量--------------")
	//必须初始化,const n1 int这种会报错
	//常量只能修饰int,fload,bool,string
	//切片数组结构体map等不行const c102 = [2]int{12, 22}会报错
	const c100 = 9
	const c101 = "lxb"
	const c102 = c100 / 3

	fmt.Println(c100, c101, c102, c103, c104, c105, c106)

	fmt.Println("------------字节大小-------------")
	//转成可以表示长度的byte切片
	data := []byte{'a', 'b'}
	dl := uint32(len(data))
	fmt.Println("dl = ", dl)
	buf := [4]byte{}
	binary.LittleEndian.PutUint32(buf[:], dl)
	fmt.Println("buf = ", buf)

	data1 := "aaaa"
	fmt.Println("data1=", []byte(data1))
	fmt.Println("-------------指针测试--------------")
	var n4 *int
	n5 := 10
	n4 = &n5
	fmt.Println("n4 指向的地址=", n4)
	fmt.Println("n4 指向的地址的值=", *n4)

	m100 := make(map[int]string, 10)
	fmt.Println("m100=", len(m100))

	m101 := []int{10, 11}
	fmt.Println("m101-1 = ", m101)

	m101 = []int{}
	fmt.Println("m101-2 = ", m101)

	var m103 []int
	m102 := []int{100, 20, 30}
	fmt.Println("m102-1", m102)
	m103 = append(m102[0:1], m102[1+1:]...)
	fmt.Println("m102-2", m103)

	// type User struct {
	// 	Name string
	// 	Age  int
	// }

	// type UserMap map[int]User

	m105 := map[int]User{
		100: {
			Name: "lxb",
			Age:  30,
		},
	}
	fmt.Println("m105= ", m105)

	bt := []int{11, 22, 33, 44, 55}
	fmt.Println("bt = ", bt[:4])

	bt2 := make([]int, 10)
	fmt.Println("bt2 len = ", len(bt2))
	fmt.Println("bt2  = ", bt2[:5])

	bt2[5] = 10
	fmt.Println(fmt.Println("bt2  = ", bt2))

	bt3 := make([][]int, 10)

	fmt.Println("bt3=", bt3)

	var bt5 int

	fmt.Println("bt5=", bt5)

	bt6 := -2 << 4
	fmt.Println("bt6=", bt6)

	bt7 := 4 >> 10
	fmt.Println("bt7=", bt7)

	bt8 := &User{
		Name: "lxb",
	}

	fmt.Println("bt8 1 = ", bt8)

	Test01(bt8)

	fmt.Println("bt8 2 = ", bt8)

	Test02(bt8)

	fmt.Println("bt8 3 = ", bt8)

	bt9 := []string{"name=lxb", "passwd=123"}
	bt10 := strings.Join(bt9, "&")
	fmt.Println("bt10 = ", bt10)

	bt11 := map[string]interface{}{
		"name":   "lxb",
		"passwd": "123",
	}

	value11 := []string{}
	for k, v := range bt11 {
		value11 = append(value11, fmt.Sprintf("%s=%s", k, v))
	}

	bt12 := strings.Join(value11, "&")
	fmt.Println("bt12 = ", bt12)

	bt13 := "+"
	bt14 := []byte(bt13)
	fmt.Printf("%T", int(bt14[0]))

	bt15 := 8 << 20
	fmt.Println("bt15 = ", bt15)

	bt16 := "abca"
	bt17 := []byte(bt16)
	fmt.Println("bt17 bytes = ", len(bt17))

	bt20 := []interface{}{1, "lxb"}
	fmt.Println("bt20 = ", bt20)

	bt19 := new(User)
	fmt.Printf("bt19 type = %T, bt19 = %v", bt19, bt19)

}

func Test01(u *User) {
	temp := u
	temp.Name = "lqm"
}

func Test02(u *User) {
	u.Name = "lyy"
}
