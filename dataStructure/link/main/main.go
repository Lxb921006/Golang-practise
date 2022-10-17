package main

import "fmt"

//单向链表
type HeroNode struct {
	No       int
	Name     string
	NickName string
	Next     *HeroNode //表示指向下一个节点
}

//给链表插入一个节点, 从尾部添加
func InsertHeroNode(head, newHeroNode *HeroNode) {
	//创建一个临时变量指向head,然后通过for循环添加到单链表尾部
	temp := head
	for {
		if temp.Next == nil { //指针类型变量在未给值之前都是一个nil
			break //当走到单链表的最后一个就退出,同时temp已经是指向到最后一个
		}
		temp = temp.Next
	}
	//将newHeroNode加入到链表的最后一个
	temp.Next = newHeroNode
}

//给链表插入一个节点, 按照编号从小到大排序(实际工作中更多的是用这个方式)
func InsertHeroNode2(head, newHeroNode *HeroNode) {
	//创建一个临时变量指向head,然后通过for循环添加到单链表尾部
	temp := head
	flag := true
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.No > newHeroNode.No {
			break
		} else if temp.Next.No == newHeroNode.No {
			//已经存在no
			fmt.Printf("no=%d已存在\n", newHeroNode.No)
			flag = false
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Printf("no=%d,已存在", newHeroNode.No)
	} else {
		newHeroNode.Next = temp.Next
		temp.Next = newHeroNode
	}

}

//显示链表的所有信息
func ListLink(head *HeroNode) {
	if head.Next == nil {
		fmt.Println("name1=", head.Name)
		return
	}
	if head.Name != "" {
		fmt.Println("name2=", head.Name)
	}
	ListLink(head.Next)
}

func ListLink2(head *HeroNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("空了")
		return
	}

	for {
		fmt.Printf("no=%d, name=%v ==>", temp.Next.No, temp.Next.Name)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
}

//删除节点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.No == id {
			flag = true
			break
		}
		temp = temp.Next
	}

	if flag {
		temp.Next = temp.Next.Next
	} else {
		fmt.Printf("no=%d不存在", id)
	}
}

//链表:是一个有序的列表,在内存存储中可能不是连续的
//应用场景:可以创建的自己的内存数据库,支持增删改查
func main() {
	//单链表
	//先创建头节点,不存储任何数据,只是为了方便链表的增删改查
	head := &HeroNode{}

	//创建一个新的节点
	h1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	h2 := &HeroNode{
		No:       2,
		Name:     "林冲",
		NickName: "豹子头",
	}
	h3 := &HeroNode{
		No:       3,
		Name:     "吴用",
		NickName: "智多星",
	}
	InsertHeroNode2(head, h2)
	InsertHeroNode2(head, h1)
	InsertHeroNode2(head, h3)
	ListLink2(head)

	DelHeroNode(head, 1)
	fmt.Println()
	ListLink2(head)
}
