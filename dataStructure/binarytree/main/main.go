package main

import "fmt"

//二叉树数据结构的遍历，前序，中序，后序遍历
//每个节点最多只有两个子节点的数据结构-二叉树结构
type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历-先输出root节点，然后再输出左子树，再输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("id=%d, name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历：先输出root的左子树，再输出root节点，再输出右子树
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("id=%d, name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历：先输出左子树，再输出右子树，再输出root节点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("id=%d, name=%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	left2 := &Hero{
		No:   5,
		Name: "鲁智深",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right3 := &Hero{
		No:   6,
		Name: "花荣",
	}

	root.Left = left1
	root.Right = right1

	right1.Right = right2

	left1.Left = left2
	left1.Right = right3

	fmt.Println("------------前序遍历-----------")
	PreOrder(root)
	fmt.Println("------------中序遍历-----------")
	InfixOrder(root)
	fmt.Println("------------后序遍历-----------")
	PostOrder(root)
}
