package main

import (
	"errors"
	"fmt"
)

//模拟环形队列-通过数组,
//核心思想就是取模(余数),队尾不能指向元素,要预留一个位置
type Queue struct {
	Maxsize int
	Array   [5]int
	Head    int //队首
	Tail    int //队尾, 不含指向的元素
}

func (q *Queue) Push(val int) (err error) {
	if q.IsFull() {

		return errors.New("队列满了")
	}

	q.Array[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.Maxsize

	return
}

func (q *Queue) Pop() (val int, err error) {
	if q.IsEmpty() {
		err = errors.New("队列空了")
		return
	}

	val = q.Array[q.Head]
	q.Head = (q.Head + 1) % q.Maxsize

	return
}

//队列是否满了
func (q *Queue) IsFull() bool {
	return (q.Tail+1)%q.Maxsize == q.Head //成立说明队列就满了
}

//队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.Tail == q.Head //成立说明队列就空了
}

//环形队列里边有多少元素
func (q *Queue) Size() int {
	return (q.Tail + q.Maxsize - q.Head) % q.Maxsize
}

//显示队列
func (q *Queue) Show() {
	size := q.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	tmep := q.Head
	for i := 0; i < size; i++ {
		fmt.Printf("q.Array[%d]=%d\n", tmep, q.Array[tmep])
		tmep = (tmep + 1) % q.Maxsize //必须要这样写,tmep++输出不了环形
	}
}

func main() {

	q := &Queue{
		Maxsize: 5,
		Head:    0,
		Tail:    0,
	}

	out := true
	key := 0
	data := 0

	for out {
		fmt.Println("\t1 添加数据")
		fmt.Println("\t2 获取数据")
		fmt.Println("\t3 队列长度")
		fmt.Println("\t4 展示队列")
		fmt.Printf("请输入选项:")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Printf("请输入要添加的数据:")
			fmt.Scanln(&data)
			err := q.Push(data)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			val, err := q.Pop()
			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Println(val)
			}
		case 3:
			val := q.Size()
			fmt.Println(val)
		case 4:
			q.Show()
		default:
			fmt.Println("输入错误")
		}
		key = 0
	}
}
