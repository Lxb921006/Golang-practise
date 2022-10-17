package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Node struct {
	Row int
	Col int
	Val int
}

func main() {
	//稀疏数组: 当一个数组中大部分元素为0或者为同一个值的数组时,可以用稀疏数组解决

	c := [5][5]int{}
	c[1][1] = 1 //1代表黑子
	c[2][3] = 2 //3代表白子

	for _, v := range c { //行
		for _, v1 := range v { //列
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}

	//遍历c, 如果有不为0,创建一个node结构体,将其放入到对应的切片即可
	var sArr []Node
	//标准的稀疏数组第一行是要将行和列的规模以及默认值都要写入到sArr
	nd := Node{
		Row: 5,
		Col: 5,
		Val: 0,
	}
	sArr = append(sArr, nd)
	for r, v := range c { //行
		for c, v1 := range v { //列
			if v1 != 0 {
				nd := Node{
					Row: r,
					Col: c,
					Val: v1,
				}
				sArr = append(sArr, nd)
			}
		}
	}

	//存盘
	file := "C:/Users/Administrator/Desktop/data.data"
	for k, v := range sArr {
		fmt.Printf("%d:%d\t %d\t %v\n", k, v.Row, v.Col, v.Val)
	}

	b, e := json.Marshal(&sArr)
	if e != nil {
		fmt.Println("序列化失败")
		return
	}

	Save(file, string(b))
	Recover(file)

	//恢复
	re := make([][]int, 5)

	for _, v := range sArr {
		if v.Val == 0 {
			for i := 0; i < v.Row; i++ {
				for t := 0; t < v.Col; t++ {
					re[i] = append(re[i], 0)
				}
			}
		} else {
			re[v.Row][v.Col] = v.Val
		}
	}

	fmt.Println("re = ", re)

}

func Save(file string, data string) {
	f, err := os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	wer := bufio.NewWriter(f)

	wer.WriteString(data + "\n")

	wer.Flush()
}

func Recover(file string) {
	f, err1 := os.OpenFile(file, os.O_RDONLY, 0777)
	var ds []Node
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		err = json.Unmarshal([]byte(str), &ds)
		if err != nil {
			fmt.Println("反序列化失败")
			return
		}
	}

	re := make([][]int, 5)

	for _, v := range ds {
		if v.Val == 0 {
			for i := 0; i < v.Row; i++ {
				for t := 0; t < v.Col; t++ {
					re[i] = append(re[i], 0)
				}
			}
		} else {
			re[v.Row][v.Col] = v.Val
		}
	}

	fmt.Println("re1 = ", re)

}
