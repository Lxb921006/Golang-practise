package main

import "fmt"

func main() {
	//切片,大致可以理解成动态数组，但是又有区别，是一种引用类型,用法跟数组差不多
	//切片定义方式1,必须make使用
	fmt.Println("-----------------切片定义方式--------------------")
	var s1 []int
	fmt.Println("s1=", s1)
	s2 := [4]int{11, 22, 33, 44}
	//切片定义方式2
	s3 := s2[1:3]
	fmt.Println("s3=", s3)
	fmt.Println("s3[0]=", s3[0])
	fmt.Println("s3 len=", len(s3))
	fmt.Println("s3 容量=", cap(s3)) //切片的容量是可以动态变化的
	s3[0] = 1212
	fmt.Println("s2=", s2)
	//切片在内存里记录了三个值，值切片的值地址，切片本身长度，容量，从底层来说是一个结构体
	// type slice struct { ptr *[4]int, len int, cap int}
	//方式3，通过make方式来创建切片，也会创建一个数组，但是这个数组是程序员不可见的，由切片在底层进行维护
	var s4 []int = make([]int, 5, 10)
	s4[0] = 5
	fmt.Println("s4=", s4)
	fmt.Println("s4 len=", len(s4))
	fmt.Println("s4 cap=", cap(s4))
	//方式4
	s5 := []int{11, 22, 33}
	fmt.Println("s5=", s5)
	fmt.Println("s5 len=", len(s5))
	fmt.Println("s5 cap=", cap(s5))
	fmt.Printf("s5 类型=%T\n", s5)
	fmt.Println("-----------------切片遍历--------------------")
	s6 := [...]int{11, 22, 33, 44}
	fmt.Println("s6=", s6)
	s7 := s6[:]
	for i := 0; i < len(s7); i++ {
		fmt.Printf("s7[%d]=%d\n", i, s7[i])
	}
	for i, v := range s7 {
		fmt.Printf("s7[%d]=%d\n", i, v)
	}
	s7[0] = 100 //这里修改也会改变s6[0]的值，因为他们都指向同一个空间
	fmt.Println("s7=", s7)
	fmt.Println("s6-2=", s6)
	fmt.Println("-----------------切片append--------------------")
	s8 := []int{11, 22, 33}
	s10 := [3]int{77, 88, 99}
	s9 := append(s8, 44, 55, 66) //可以追加多个， 切片类型才能append
	s8 = append(s8, 44, 55, 66)
	s8 = append(s8, s10[:]...) //可以追加切片
	fmt.Println("s8=", s8)
	fmt.Println("s8 len=", len(s8))
	fmt.Println("s8 cap=", cap(s8))
	fmt.Println("s9=", s9)
	fmt.Println("s9 len=", len(s9))
	fmt.Println("s9 cap=", cap(s9))
	fmt.Println("-----------------切片copy--------------------")
	s11 := []int{11, 22, 33}
	s12 := make([]int, 6)
	copy(s12, s11) //切片类型才能拷贝
	fmt.Println("s11=", s11)
	fmt.Println("s12=", s12)
	s11[0] = 110 //s11，s12是两个独立的空间，修改元素的值互不影响
	fmt.Println("s11-2=", s11)
	fmt.Println("s12-2=", s12)
	fmt.Println("-----------------切片案例--------------------")
	s13 := make([]int, 1)
	s14 := []int{11, 22, 33}
	copy(s13, s14)
	fmt.Println("s13=", s13)
	test01(s13)
	fmt.Println("s13=", s13)
	fmt.Println("-----------------切片string-------------------")
	s15 := "lxb"       //字符串在底层其实是一个byte数组存在，可以进行切片处理，string是不可以变的，s15[0] = 'g',这样会报错，如果真要改只能先将s15转成[]byte切片类型，然后s[0]=o修改再转成string(s15)
	s18 := []rune(s15) //需要注意的是：可以处理英文跟数字，但是不能处理中文，因为byte是按字节编码的，中文占三个字节会提示超出码值范围编译不通过，解决方法是将中文字符串转成[]rune（它兼容中文）即可
	s18[0] = '廖'
	fmt.Println("s18", string(s18))
}

func test01(n []int) { //切片是引用类型，传入的是地址，地址指向的空间跟main里的s13指向的空间是同一个，因此在函数里修改s13[0]的值，在main里边的s13[0]也是会变化
	n[0] = 100
}
