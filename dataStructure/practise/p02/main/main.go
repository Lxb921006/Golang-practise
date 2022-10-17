package main

import "fmt"

//用链表模拟队列
type RingLink struct {
	No   int
	Name string
	Next *RingLink
}

//向环形链表添加数据
func AddData(head *RingLink, new *RingLink) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = new
}

func GetData(head *RingLink) {
	fmt.Println("----------------获取数据---------------")
	temp := head
	if temp.Next == nil {
		fmt.Println("空了")
	} else {
		fmt.Printf("获取No=%d\n", temp.Next.No)
		temp.Next = temp.Next.Next
	}
	fmt.Println("----------------获取数据---------------")
}

func Show(head *RingLink) {
	fmt.Println("----------------展示所有数据---------------")
	temp := head
	if temp.Next == nil {
		fmt.Println("空了")
		return
	}

	for {
		fmt.Printf("展示No=%d\n", temp.Next.No)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
	fmt.Println("----------------展示所有数据---------------")
}

func main() {
	head := &RingLink{}
	id := 0
	out := true
	data := 0
	for out {
		fmt.Println("1 添加数据")
		fmt.Println("2 获取数据")
		fmt.Println("3 查看数据")
		fmt.Println("输入1或者2或者3")
		fmt.Scanln(&id)
		switch id {
		case 1:
			fmt.Println("请输入:")
			fmt.Scanln(&data)
			nd := &RingLink{
				No: data,
			}
			AddData(head, nd)
		case 2:
			GetData(head)
		case 3:
			Show(head)
		default:
			fmt.Println("输入错误")
		}
		id = 0
	}
}
