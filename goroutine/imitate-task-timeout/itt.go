package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var lock sync.Mutex
	var taskCh = make(chan int)
	var done = make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			go func() {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						return
					case v, ok := <-taskCh:
						if !ok {
							return
						}
						task(v)
						done <- 1
					}
				}
			}()

			select {
			case <-done:
				lock.Lock()
				counter++
				lock.Unlock()
			case <-time.After(time.Second * 3):
				fmt.Println("time out")
			}
		}()
	}

	go func() {
		for i := 0; i < 50; i++ {
			taskCh <- i
		}

		close(taskCh)
	}()

	wg.Wait()

	fmt.Printf("总共完整了: %d个任务", counter)
}

func task(i int) {
	time.Sleep(time.Duration(rand.Intn(30)+1) * time.Second)
	fmt.Printf("task %d done\n", i)
}
