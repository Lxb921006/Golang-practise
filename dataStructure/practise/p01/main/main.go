package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

type Queue struct {
	Maxsize int
	Array   [5]int
	Head    int //队首
	Tail    int //队尾, 不含指向的元素
}

func (q *Queue) Push() {
	for {
		rd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3)
		val := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000)
		if q.IsFull() {
			fmt.Println("服务窗口满人了")
		} else {
			lock.Lock()
			q.Array[q.Tail] = val
			q.Tail = (q.Tail + 1) % q.Maxsize
			lock.Unlock()
		}
		time.Sleep(time.Duration(rd) * time.Second)
	}
}

func (q *Queue) Pop(g int) {
	for {
		rd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3)
		if q.IsEmpty() {
			fmt.Println("没人了")
		} else {
			lock.Lock()
			val := q.Array[q.Head]
			fmt.Printf("%d号协程服务了-->%v号客户\n", g, val)
			q.Head = (q.Head + 1) % q.Maxsize
			lock.Unlock()
		}
		time.Sleep(time.Duration(rd) * time.Second)
	}
}

//队列是否满了
func (q *Queue) IsFull() bool {
	return (q.Tail+1)%q.Maxsize == q.Head //成立说明队列就满了
}

//队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.Tail == q.Head //成立说明队列就空了
}

func main() {
	q := &Queue{
		Maxsize: 5,
		Head:    0,
		Tail:    0,
	}

	go q.Push()

	for i := 0; i < 2; i++ {
		go q.Pop(i)
	}

	for 1 > 0 {
		continue
	}

}
