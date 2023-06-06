package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Task func(cxt context.Context)

type MultiWork struct {
	Works chan Task
	Limit chan struct{}
	Wg    sync.WaitGroup
	Max   int
}

func NewMultiWork(workers int) *MultiWork {
	nm := &MultiWork{
		Works: make(chan Task),
		Limit: make(chan struct{}, workers),
	}

	go func() {
		for task := range nm.Works {

			nm.Limit <- struct{}{}
			nm.Wg.Add(1)

			fmt.Println(runtime.NumGoroutine())

			go func(task Task) {
				defer func() {
					nm.Wg.Done()
					<-nm.Limit
				}()

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				var done = make(chan struct{})

				go func() {
					task(ctx)
					done <- struct{}{}
				}()

				select {
				case <-done:
				case <-ctx.Done():
					fmt.Println("timeout")
				}

			}(task)
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

func main() {

	rand.Seed(time.Now().UnixNano())

	nm := NewMultiWork(10)

	for i := 1; i < 100; i++ {
		nm.Add(func(cxt context.Context) {

			select {
			case <-time.After(time.Second * time.Duration(1+rand.Intn(5))):
				fmt.Println(rand.Intn(1000))
			case <-cxt.Done():
			}

		})
	}

	nm.Close()
}
