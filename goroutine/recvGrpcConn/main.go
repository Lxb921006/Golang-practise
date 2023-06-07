package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Task func(ctx context.Context)

type MultiWork struct {
	Works   chan Task
	Limit   chan struct{}
	Wg      sync.WaitGroup
	lock    sync.Mutex
	running int
}

func NewMultiWork(workers int) *MultiWork {
	nm := &MultiWork{
		Works: make(chan Task),
		Limit: make(chan struct{}, workers),
	}

	go func() {
		for task := range nm.Works {
			//fmt.Println("running = ", nm.running)
			fmt.Println("gn = ", runtime.NumGoroutine())

			nm.Limit <- struct{}{}
			nm.Wg.Add(1)

			go func(task Task) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer func() { nm.dec(); nm.Wg.Done() }()
				defer cancel()

				var done = make(chan struct{})

				go func() {
					task(ctx)
					done <- struct{}{}
				}()

				select {
				case <-ctx.Done():
					fmt.Println("timeout")
				case <-done:
				}

				<-nm.Limit
			}(task)
			nm.inc()
		}
	}()

	return nm
}

func (mw *MultiWork) Add(task Task) {
	mw.Works <- task
}

func (mw *MultiWork) Close() {
	mw.Wg.Wait()
	close(mw.Works)
}

func (mw *MultiWork) inc() {
	mw.lock.Lock()
	mw.running++
	mw.lock.Unlock()
}

func (mw *MultiWork) dec() {
	mw.lock.Lock()
	mw.running--
	mw.lock.Unlock()
}

func main() {

	rand.Seed(time.Now().UnixNano())

	nm := NewMultiWork(10)

	for i := 1; i < 100; i++ {
		nm.Add(func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
				default:
					if err := ctx.Err(); err != nil {
						return
					}
					time.Sleep(time.Second * time.Duration(1+rand.Intn(5)))
					fmt.Println(rand.Intn(1000))
				}
			}

		})
	}

	nm.Close()
	i := 25

	for i >= 0 {
		fmt.Println("close gn = ", runtime.NumGoroutine())
		fmt.Println("running = ", nm.running)
		i--
		time.Sleep(time.Second)
	}
}
