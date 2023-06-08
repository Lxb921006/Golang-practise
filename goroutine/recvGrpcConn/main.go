package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Task func()

type MultiWork struct {
	Works chan Task
	Limit chan struct{}
	Wg    sync.WaitGroup
	lock  sync.Mutex
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
			go func(task Task) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				done := make(chan struct{})
				go func() {
					task()
					close(done)
				}()

				select {
				case <-done:
					// task completed successfully
				case <-ctx.Done():
					fmt.Println("task canceled:", ctx.Err())
				}

				<-nm.Limit

			}(task)

		}
	}()

	return nm
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nm := NewMultiWork(10)

	task := func() {
		defer nm.Wg.Done()
		time.Sleep(time.Second * time.Duration(rand.Intn(1000)+1))
		fmt.Println("task finished ", rand.Intn(1000))
	}

	for range [100]struct{}{} {
		nm.Works <- task
	}

	nm.Wg.Wait()
	close(nm.Works)

	i := 25
	for i > 0 {
		fmt.Println("gn = ", runtime.NumGoroutine())
		i--
		time.Sleep(time.Second * 1)
	}

}
