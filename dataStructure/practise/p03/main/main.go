package main

import (
	"fmt"
)

//环形单向链表解决约瑟夫问题
type JosepuhIssue struct {
	No   int
	Next *JosepuhIssue
}

func AddData(n int) *JosepuhIssue {
	first := &JosepuhIssue{}
	cur := &JosepuhIssue{}

	if n < 1 {
		return first
	}

	for i := 1; i <= n; i++ {
		nj := &JosepuhIssue{
			No: i,
		}
		if i == 1 {
			first = nj       //fisrt添加数据之后就先不去动它
			cur = nj         //辅助节点
			cur.Next = first //这里是构成环形
		} else {
			cur.Next = nj
			cur = nj //辅助节点,保存的是每次循环的nj
			nj.Next = first
		}

	}

	return first
}

func GetData(first *JosepuhIssue, startNo, countNo int) {
	if first.Next == nil {
		fmt.Println("空了")
	}

	s := func() (n int) {
		temp2 := first
		for {
			n++
			if temp2.Next == first {
				break
			}
			temp2 = temp2.Next
		}
		return
	}()

	if startNo > s && startNo < 0 {
		fmt.Println("超出范围")
		return
	}

	//先找到起始位置
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
	}

	temp3 := first //记录开始前的k位置==startNo

	//移动多少步
	for {
		for i := 1; i <= countNo-1; i++ {
			temp3 = first //保存上一个startNo的位置
			first = first.Next
		}
		fmt.Printf("编号%d出列\n", first.No)
		first = first.Next //移动到哪个节点就删除哪个节点
		temp3.Next = first //将保存的startNo位置重新指向被删除节点的下一个节点
		if temp3 == first {
			break
		}

	}
	fmt.Printf("最后一个编号%d出列\n", first.No)
}

func Show(first *JosepuhIssue) {
	cur := first
	for {
		fmt.Println("data = ", cur.Next)
		if cur.Next == first {
			break
		}
		cur = cur.Next
	}
}

func main() {
	f := AddData(500)
	GetData(f, 20, 31)
	// Show(f)
}
