package main

import (
	"fmt"
	"runtime"
	"time"
)

type Pool struct {
	Work chan func()
	Sem  chan bool //有缓冲的chan，限制goroutine number
}

func (p *Pool) Worker(task func()) {
	defer func() { <-p.Sem }()
	for {
		task()
		// <-p.Work

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
	pool := NewPool(2)
	for i := 0; i < 20; i++ {
		pool.Task(func() {
			time.Sleep(time.Second)
			fmt.Printf("goroutine num = %d, time = %v\n", runtime.NumGoroutine(), time.Now())
		})
	}
	// time.Sleep(time.Second * 2)
}
