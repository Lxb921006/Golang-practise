package main

import (
	"fmt"
)

//环形链表(对应内置库的ring):头节点必须有值
type RingLink struct {
	No   int
	Name string
	Next *RingLink
}

//向环形链表添加数据
func InsertData(head *RingLink, new *RingLink) {
	//判断是否是第一个添加进去的数据
	if head.Next == nil {
		head.No = new.No
		head.Name = new.Name
		head.Next = head //目的是需要第一个添加进去的数据形成环状,此时head已经指向one并且已经初始化, 所以head.Next已经是one
		return
	}

	//找到环形的最后一个节点
	temp := head
	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}

	//加入到环形队列
	temp.Next = new
	new.Next = head
}

func Delete(head *RingLink, id int) *RingLink {
	//环形链表空的情况
	if head.Next == nil {
		fmt.Println("空")
		return head
	}

	temp := head
	flag := false

	//环形链表只有一个数据的情况
	if temp.Next == head {
		temp.Next = nil
		return temp
	}

	//环形链表里有两个以上数据的情况
	for {
		if temp.Next.No == id {
			break
		}
		if temp.Next == head {
			flag = true
			break
		}
		temp = temp.Next
	}

	if flag {
		fmt.Println("没找到")
		return head
	}

	if temp.Next != head {
		temp.Next = temp.Next.Next
	} else {
		temp.Next = head.Next
		head = temp
		fmt.Println("head = ", head)
	}

	return head
}

func Show(head *RingLink) {
	fmt.Println("----------输出如下-----------")
	temp := head
	if head.Next == nil {
		fmt.Println("空")
		return
	}

	for {
		fmt.Println("data = ", temp.Next)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
}

func main() {
	//环形链表需要给头节点放入数据
	head := &RingLink{}
	one := &RingLink{
		No:   1,
		Name: "one",
	}
	two := &RingLink{
		No:   2,
		Name: "two",
	}
	// three := &RingLink{
	// 	No:   3,
	// 	Name: "three",
	// }
	// four := &RingLink{
	// 	No:   4,
	// 	Name: "four",
	// }
	InsertData(head, one)
	InsertData(head, two)
	// InsertData(head, three)
	// InsertData(head, four)
	head = Delete(head, 1)
	Show(head)
}
