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
	Works chan Task
	Limit chan struct{}
	Wg    sync.WaitGroup
	lock  sync.Mutex
	done  chan struct{}
}

func NewMultiWork(workers int) *MultiWork {
	nm := &MultiWork{
		Works: make(chan Task),
		Limit: make(chan struct{}, workers),
		done:  make(chan struct{}, 1),
	}

	go func() {
		for task := range nm.Works {
			nm.Limit <- struct{}{}

			nm.Wg.Add(1)
			go func(task Task) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				go func() {
					select {
					case <-ctx.Done():
						return
					case <-nm.done:
						return
					}
				}()
				task(ctx)
				nm.done <- struct{}{}
				<-nm.Limit
			}(task)

		}
	}()

	return nm
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nm := NewMultiWork(10)

	task := func(ctx context.Context) {
		defer nm.Wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("task cancel")
				return
			default:
				if err := ctx.Err(); err != nil {
					return
				}
				time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
				fmt.Println("task finished ", rand.Intn(1000))
				return
			}
		}

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
