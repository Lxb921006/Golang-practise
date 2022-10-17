package main

import (
	"errors"
	"fmt"
)

//模拟队列-先进先出
type Queue struct {
	MaxSize int
	Array   [5]int
	Front   int //队首,但不包含指向的位置, 根据输出不断的变化
	Rear    int //队尾, 包含指向的位置, 根据输入不断的变化
}

func (q *Queue) AddData(val int) (err error) {
	if q.Rear == q.MaxSize-1 {
		fmt.Println("q.Front = ", q.Front)
		return errors.New("队列满了")
	}

	q.Rear++ //如果有新的数据,就往后移

	q.Array[q.Rear] = val

	return
}

func (q *Queue) GetData() (val int, err error) {
	if q.Front == q.Rear {
		fmt.Println("q.Rear = ", q.Rear)
		err = errors.New("队列空了,清先添加")
		return
	}

	q.Front++

	val = q.Array[q.Front]

	return
}

//显示队列, 从队首遍历到队尾
func (q *Queue) Showqueue() {
	if q.Rear < 0 {
		fmt.Println("请先向队列添加数据")
		return
	}

	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("q.Array[%d]=%d\n", i, q.Array[i])
	}
}

func main() {
	//队列：先进先出，实现方式：数组或者链表
	//队列实现-数组非环形
	q := &Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}

	out := true
	key := 0
	data := 0

	for out {
		fmt.Println("\t1 添加数据")
		fmt.Println("\t2 获取数据")
		fmt.Println("\t3 展示队列")
		fmt.Printf("请输入选项:")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Printf("请输入要添加的数据:")
			fmt.Scanln(&data)
			err := q.AddData(data)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			val, err := q.GetData()
			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Println(val)
			}
		case 3:
			q.Showqueue()
		default:
			fmt.Println("输入错误")
		}
		key = 0
	}
}
