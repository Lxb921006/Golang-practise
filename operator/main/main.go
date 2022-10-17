package main

import (
	"fmt"
)

//算术运算符
func main() {
	a, b := 0, 3
	a++
	b--
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println(10.5 / 4) //Golang里边两边都是整数相除只保留整数结果，小数点后面会去掉
	//需要保留后面小数可以如下设置
	fmt.Println(10.0 / 4)
	//取模的公式：a % b = a - a / b * b
	fmt.Println(10 % 4)   // 2
	fmt.Println(-10 % 4)  // -10 - (-10) / 4 * 4 = -10 - (-8) = -2
	fmt.Println(10 % -4)  // 10 - 10 / 4 * 4 = 2
	fmt.Println(-10 % -4) // -10 - (-10) / -4 * -4 = -10 - (-8) = -2
	//自增，自减注意事项
	//不可以直接这样赋值a = 10;b := a++或者b := a--或者i=a++,i=a--都不可以，必须独立使用，如a++,a--
	day := 97 % 7
	week := 97 / 7
	fmt.Println("day=", day)
	fmt.Println("week=", week)

	dd1 := "lxb"
	dd2 := "lxb"
	dd3 := 3 > 2
	fmt.Println("dd1==dd2", dd1 == dd2)
	fmt.Println("dd3==", dd3)
	fmt.Println("dd2.len==", len(dd2))
	fmt.Printf("dd2[1]==%q\n", dd2[1])
	//逻辑
	if dd1 != "" && dd2 == "lxb1" {
		fmt.Println("ok")
	}
	if dd1 != "" || dd2 == "lxb1" {
		fmt.Println("ok")
	}
	if !(dd1 == dd2) || dd3 {
		fmt.Println("result is bingo.")
	} else {
		fmt.Println("result is error.")
	}
	if dd3 {
		fmt.Println("dd3 is ture.")
	} else {
		fmt.Println("dd3 is fale.")
	}
	dd4 := 100
	if dd4 > 10 {
		fmt.Println(dd4)
	} else if dd4 == 100 {
		fmt.Println(dd4)
	} else {
		fmt.Println("false")
	}
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	single := "lxbb"
	switch single {
	case "lxb":
		fmt.Println("myself")
	case "lqm":
		fmt.Println("wife")
	case "lyy":
		fmt.Println("sister")
	default:
		fmt.Println("nobody")
	}
	if dd1 == "lxb" && test() {
		fmt.Println("调用test()比较结果为true")
	} else {
		fmt.Println("调用test()比较结果为false")
	}
	//赋值运算
	uu1 := 10
	uu2 := 10
	// uu1 += uu2
	// uu1 -= uu2
	// uu1 *= uu2
	// uu1 /= uu2
	uu1 %= uu2
	fmt.Println("uu1=", uu1)

	uu4 := test2() + 10
	fmt.Println("uu4=", uu4)
	//左移运算符，跟二进制有关
	uu3 := 10
	uu3 <<= 2
	fmt.Println("uu3=", uu3)

	//面试题
	var uu5 = 1
	var uu6 = 2
	fmt.Printf("交换前uu5=%d, uu6=%d\n", uu5, uu6)
	uu5, uu6 = uu6, uu5
	fmt.Printf("交换后uu5=%d, uu6=%d\n", uu5, uu6)

	//位运算符
	fmt.Println(3 & 2)   //2
	fmt.Println(3 | 2)   //3
	fmt.Println(3 ^ 2)   //1  按位异或，两位一个为0一个为1，结果为1，否则为0
	fmt.Println(-3 & 2)  //0
	fmt.Println(-6 | 7)  //-1
	fmt.Println(-8 | 7)  //-1
	fmt.Println(-10 & 8) //0  //按位与，两位全为1，结果为1，否则为0
	fmt.Println(-7 | 6)  //-1 //按位或，两位有一个为1，结果为1，否则为0
	fmt.Println(4 &^ 3)  //4
	fmt.Println(-4 &^ 3) //4  &^ 运算：位清空运算，和被运算变量位置有关系。计算x&^y 如果ybit位上的数是0则取x上对应位置的值， 如果ybit位上为1则取结果位上取0

	//<< 运算：a << b就表示把a转为二进制后左移b位（在后面添b个0）。
	//>> 运算：a >> b表示二进制右移b位（去掉末b位），相当于a除以2的b次方（取整）
	//移位运算
	//右移运算符 >>:低位溢出，符号位不变，并用符号位补溢出的高位
	//左移运算符 <<:符号位不变，低位补0
	//只要是运算都是以补码的方式进行
	h1 := 1 >> 2
	h2 := 1 << 2
	h3 := 6 >> 2
	h4 := 6 << 2
	h5 := 7 << 3
	h6 := 7 >> 3
	h7 := 7 >> 2
	h8 := -7 >> 2
	h9 := -8 >> 3
	h10 := -7 << 2
	h11 := -7 << 6
	h12 := 10 >> 6
	h13 := -10 << 3
	h14 := 4 >> 2
	h15 := 7 >> 2
	h16 := 1 << 16
	fmt.Println("h1=", h1)
	fmt.Println("h2=", h2)
	fmt.Println("h3=", h3)
	fmt.Println("h4=", h4)
	fmt.Println("h5=", h5)
	fmt.Println("h6=", h6)
	fmt.Println("h7=", h7)
	fmt.Println("h8=", h8)
	fmt.Println("h9=", h9)
	fmt.Println("h10=", h10)
	fmt.Println("h11=", h11)
	fmt.Println("h12=", h12)
	fmt.Println("h13=", h13)
	fmt.Println("h14=", h14)
	fmt.Println("h15=", h15)
	fmt.Println("h16=", h16)
}

func test2() int {
	return 100
}

func test() bool {
	return true
}
