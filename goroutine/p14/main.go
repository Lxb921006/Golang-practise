package main

import (
	"fmt"
	"time"
)

type Pool struct {
	Work chan func()
	Sem  chan bool
}

func (p *Pool) Worker(task func()) {
	defer func() { <-p.Sem }()
	for {
		task()
		task = <-p.Work
	}
}

func (p *Pool) Task(task func()) {
	select {
	case p.Work <- task:
	case p.Sem <- true:
		go p.Worker(task)
	}
}

func NewPool(size int) *Pool {
	return &Pool{
		Work: make(chan func()),
		Sem:  make(chan bool, size),
	}
}

func main() {
	pool := NewPool(20)
	for i := 0; i < 100; i++ {
		pool.Task(func() {
			time.Sleep(time.Second)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		})
	}

	time.Sleep(time.Second * 10)
}
