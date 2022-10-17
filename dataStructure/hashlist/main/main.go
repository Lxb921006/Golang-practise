package main

import "fmt"

//员工信息-单链表
type Emp struct {
	Id   int
	Name string
	Next *Emp //后面添加的Emp都加到每一个Emp.Next里
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表 %d, 雇员信息 %v\n", e.Id%7, e)
}

//不带表头，第一个节点就存放Emp
type EmpLink struct {
	Head *Emp //存的是第一个Emp
}

//添加时还要保证Id从小到大添加
func (el *EmpLink) Add(e *Emp) {
	if el.Head == nil {
		el.Head = e
	} else {
		temp := el.Head
		temp2 := el.Head //这个是保存上一个Emp
		for {
			if temp == nil {
				break
			}
			//Emp.id从小到大排序
			if temp.Id >= e.Id {
				break
			}
			temp2 = temp
			temp = temp.Next
		}
		if temp != nil && temp.Id == temp2.Id && temp.Id != e.Id {
			el.Head = e
			el.Head.Next = temp2
			temp2.Next = nil
		} else {
			e.Next = temp2.Next
			temp2.Next = e
		}

	}
}

//删除雇员
func (el *EmpLink) Delete(no int) (emp *Emp) {
	temp := el.Head
	temp2 := el.Head

	if temp == nil {
		fmt.Printf("该雇员%d不存在\n", no)
		return
	}

	for {
		if temp == nil {
			break
		}

		if temp.Id == no {
			break
		}

		temp2 = temp
		temp = temp.Next
	}

	if temp != nil && temp.Id == temp2.Id {
		el.Head = temp.Next
	} else if temp != nil {
		temp2.Next = temp.Next
	} else {
		fmt.Printf("该雇员%d没有找到\n", no)
	}
	return
}

//修改雇员信息
func (el *EmpLink) Edit(no int, name string) (emp *Emp) {
	temp := el.Head

	if temp == nil {
		fmt.Printf("该雇员%d不存在\n", no)
		return
	}

	for {
		if temp == nil {
			break
		}
		if temp.Id == no {
			temp.Name = name
			break
		}
		temp = temp.Next
	}

	return
}

func (el *EmpLink) Find(no int) (emp *Emp) {
	temp := el.Head
	for {
		if temp == nil {
			break
		}
		if temp.Id == no {
			return temp
		}
		temp = temp.Next
	}

	return
}

//显示链表信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表 %d 为空\n", no)
		return
	}

	temp := el.Head

	for {
		if temp == nil {
			break
		}
		fmt.Printf("链表 %d, 雇员id=%d->, 雇员名字---->%s", no, temp.Id, temp.Name)
		temp = temp.Next
	}
	fmt.Println()
}

//哈希列表(散列)
type HashTable struct {
	LinkArr [7]EmpLink //EmpLink类型的数组也就是数组里存的是EmpLink结构体
}

func (h *HashTable) Add(e *Emp) {
	//使用散列函数，确定将该雇员添加到那个链表
	linkNo := h.HashFunc(e.Id)
	h.LinkArr[linkNo].Add(e)
}

func (h *HashTable) Delete(no int) {
	//使用散列函数，确定将该雇员添加到那个链表
	linkNo := h.HashFunc(no)
	h.LinkArr[linkNo].Delete(no)
}

func (h *HashTable) Edit(no int, name string) {
	//使用散列函数，确定将该雇员添加到那个链表
	linkNo := h.HashFunc(no)
	h.LinkArr[linkNo].Edit(no, name)
}

func (h *HashTable) Show() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

//通过id找出雇员信息
func (h *HashTable) Find(no int) (emp *Emp) {
	//使用散列函数找出该no是在哪个链表里
	linkNo := h.HashFunc(no)
	emp = h.LinkArr[linkNo].Find(no)
	return
}

//这个是要把雇员id添加到数组的哪个位置(下标)里,通过取模的方式来确定
func (h *HashTable) HashFunc(id int) int {
	return id % 7 //对应链表的下标
}

func main() {
	h := &HashTable{}

	e1 := &Emp{
		Id:   30,
		Name: "lxb",
	}
	e2 := &Emp{
		Id:   51,
		Name: "lqm",
	}
	e3 := &Emp{
		Id:   44,
		Name: "lyy",
	}
	e4 := &Emp{
		Id:   10,
		Name: "lch",
	}
	e5 := &Emp{
		Id:   17,
		Name: "llt",
	}
	e6 := &Emp{
		Id:   37,
		Name: "lll",
	}
	e7 := &Emp{
		Id:   17,
		Name: "lxx",
	}

	h.Show()
	fmt.Println("----------添加雇员信息------------")
	h.Add(e6)
	h.Add(e1)
	h.Add(e2)
	h.Add(e3)
	h.Add(e4)
	h.Add(e5)
	h.Add(e7)
	fmt.Println("----------查看所有雇员信息------------")
	h.Show()
	fmt.Println("----------查看指定雇员信息------------")
	emp := h.Find(30)
	if emp != nil {
		emp.ShowMe()
	} else {
		fmt.Println("没找到")
	}
	fmt.Println("----------删除指定雇员------------")
	h.Delete(30)
	h.Show()
	fmt.Println("----------修改指定雇员信息------------")
	h.Edit(44, "lxb222222")
	h.Show()
}
