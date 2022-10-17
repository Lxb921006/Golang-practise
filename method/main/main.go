package main

import "fmt"

type M1 struct {
	Name string
}

type M2 struct { //正常开发中，不建议这样命名，命名要清晰明目
	Num float64
}

type Inter int //不只是struct类型可以绑定方法，其他的类型也可以

type Cal struct {
	Num1 float64
	Num2 float64
}

//给M1结构体类型绑定一个test方法.
func (m M1) test() {
	m.Name = "lqm" //结构体是值类型，外部的修改不会影响这里的值
	fmt.Println("test()=", m.Name)
}

func (m M1) test02() {
	fmt.Printf("%v是我wife\n", m.Name)
}

func (m M1) jisuan() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Printf("%v计算的res=%d\n", m.Name, res)
}

func (m M1) jisuan02(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Printf("%v计算的res=%d\n", m.Name, res)
}

//注意这里的m是可用可不用，不用也不会报错
func (m M1) jisuan03(n int) (string, int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return m.Name, res
}

//要注意，这里的m跟m3是不一样的，在golang里，这里的m是一个新的变量或者是新数据类型，这里的方式实参都是通过结构体M2值拷贝实现
func (m M2) mianji() float64 {
	return 3.14 * m.Num
}

//通常情况下都用以下的方式，说明传入的m是一个地址，而这个地址指向的就是M2这个结构体，这里的m也就等同于m3，效率比较高
func (m *M2) mianji02() float64 {
	m.Num = 4.2 //这里会改变main()中的m.Num的值，因为这里获取的地址是指向M2这个结构体
	fmt.Printf("mianji02() m3=%p\n", m)
	return 3.14 * m.Num //这里的m.Num实际上底层还是(*m).Num，只是在golang底层优化后可写成m.Num
}

func (i Inter) Print() { //改成大写Print
	fmt.Println("i=", i)
}

func (i *Inter) change() {
	*i += 1
}

func (m *M1) String() string {
	res := fmt.Sprintf("Name=[%v]", (*m).Name)
	return res
}

func (m *M1) exercise01() {
	for i := 1; i <= 10; i++ {
		for t := 1; t <= 8; t++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (m *M1) exercise02(n1, n2 int) {
	for i := 1; i <= n1; i++ {
		for t := 1; t <= n2; t++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (m *M1) exercise03(n1, n2 int) int {
	return n1 * n2
}

func (m *M1) exercise04(n int) string {
	if n%2 == 0 {
		return "偶数"
	}
	return "奇数"
}

func (c *Cal) exercise05(n3 byte) (res float64) {
	switch n3 {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '/':
		res = c.Num1 / c.Num2
	case '*':
		res = c.Num1 * c.Num2
	}
	return
}

func (c *Cal) exercise06(n int) {
	for i := 1; i <= n; i++ {
		for t := 1; t <= i; t++ {
			fmt.Printf("%d * %d = %d\t", t, i, i*t)
		}
		fmt.Println()
	}
}

func (c *Cal) exercise07() {
	n := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(n); i++ {
		for t := 0; t < len(n[i]); t++ {
			if i == t {
				for y := 0; y < len(n); y++ {
					fmt.Printf("%d ", n[y][t])
				}
				fmt.Println()
			}
		}

	}
}

//test方法和M1类型绑定
//m1会传给test()方法中的m
//test方法只能是M1的实例调用也就是如下例子的m1，不能使用其他类型来调用
func main() {
	m1 := M1{}
	m1.Name = "lqm"
	m1.test()
	fmt.Println("main()=", m1.Name)
	m1.test02()
	(&m1).jisuan()
	m1.jisuan02(10)
	name, m2 := m1.jisuan03(10)
	fmt.Println("main() m2=", name, m2)
	m3 := M2{}
	fmt.Printf("main() m3=%p\n", &m3)
	m3.Num = 2.2
	m4 := m3.mianji()
	m5 := m3.mianji02()
	//标准的方式应该是，但是在golang底层已经优化过，所以也可以这样写m6 := m3.mianji02()，但是底层还是m6 := (&m3).mianji02()
	m6 := (&m3).mianji02()
	fmt.Println("main() m3.Num=", m3.Num)
	fmt.Println("main() m4=", m4)
	fmt.Println("main() m5=", m5)
	fmt.Println("main() m6=", m6)
	var m7 Inter = 10
	(&m7).Print()
	m7.change()
	fmt.Println("m7=", m7)
	fmt.Println(&m1) //如果一个类型实现了String()方法，那么fmt.Println会默认调用String()进行输出
	fmt.Println("----------------------练习-----------------------")
	cal := Cal{}
	cal.Num1 = 100
	cal.Num2 = 1
	(&m1).exercise01()
	m1.exercise02(5, 5)
	m8 := m1.exercise03(2, 2)
	fmt.Println("main() m8=", m8)
	m9 := (&m1).exercise04(19)
	fmt.Println("main() m9=", m9)
	m10 := (&cal).exercise05('*')
	fmt.Println("main() m10=", m10)
	cal.exercise06(9)
	cal.exercise07()
	fmt.Println("----------------------函数跟方法的区别-----------------------")
	//Test07函数接受的类型是M1结构体，所以传递的值也只能是M1实例，不能是其他类型，地址也不可以，如Test07(&m1)会报错无法编译通过，但是方法是可以的
	//函数定义好的接受参数类型就需要接受什么样的参数类型，就不能使用其他的类型，但是方法可以
	Test07(m1)
	(&m1).jisuan() //这里虽然传入的是地址，但是因为m类型是M1这个结构体,并不是指针类型，所以这里还是通过值拷贝来输出m.Name，跟外部的m.Name的值互不影响
}

func Test07(m M1) {
	fmt.Println("Test07=", m.Name)
}
