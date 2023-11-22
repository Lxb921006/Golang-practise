package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//var to bool
//var counter int
//var lock sync.Mutex

type result struct {
	resp string
	id   int
	ctx  context.Context
}

type pool struct {
	workers int
	wg      *sync.WaitGroup
	lock    *sync.Mutex
	once    *sync.Once
	taskCh  chan func() result
	ctx     context.Context
	done    chan struct{}
}

func newPool(w int, ctx context.Context) *pool {
	return &pool{
		workers: w,
		taskCh:  make(chan func() result),
		done:    make(chan struct{}, 1),
		ctx:     ctx,
		wg:      new(sync.WaitGroup),
		once:    new(sync.Once),
		lock:    new(sync.Mutex),
	}
}

func (p *pool) start() *sync.WaitGroup {
	p.wg.Add(p.workers)
	for i := 0; i < p.workers; i++ {
		go p.work()
	}

	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			case <-p.done:
			case <-time.After(time.Second * 3):
				fmt.Println("time out")
			}
		}
	}()

	return p.wg
}

func (p *pool) work() {
	defer p.wg.Done()
	for v := range p.taskCh {
		resp := v()
		fmt.Println(resp)
		p.done <- struct{}{}
	}
}

func (p *pool) stop() {
	p.once.Do(func() {
		close(p.taskCh)
	})
}

func (p *pool) wait() {
	p.wg.Wait()
}

func (p *pool) addTask(task func() result) {

	p.taskCh <- task

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newPool(5, ctx)
	job := p.start()

	go func() {
		for i := 0; i < 50; i++ {
			resp := result{
				id: i,
			}
			p.addTask(func() result {
				time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
				return resp
			})
		}
		p.stop()
	}()

	job.Wait()

	//var workers = 5
	//var wg sync.WaitGroup
	//var taskCh = make(chan int, 1)
	//var done = make(chan int)
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//wg.Add(workers)
	//for i := 0; i < workers; i++ {
	//	go func() {
	//		go func() {
	//			defer wg.Done()
	//			for {
	//				select {
	//				case <-ctx.Done():
	//					return
	//				case v, ok := <-taskCh:
	//					if !ok {
	//						return
	//					}
	//
	//					task(v)
	//					done <- 1
	//				}
	//			}
	//		}()
	//	}()
	//}
	//
	//go func() {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case <-done:
	//			lock.Lock()
	//			to = false
	//			lock.Unlock()
	//		case <-time.After(time.Second * 3):
	//			lock.Lock()
	//			to = true
	//			lock.Unlock()
	//			fmt.Println("time out ", runtime.NumGoroutine())
	//		}
	//	}
	//}()
	//
	//go func() {
	//	for i := 0; i < 50; i++ {
	//		taskCh <- i
	//	}
	//	close(taskCh)
	//}()
	//
	//wg.Wait()
	//
	//fmt.Printf("总共完整了: %d个任务", counter)
}

//func task(i int) {
//	time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
//	lock.Lock()
//	if !to {
//		fmt.Printf("task %d done\n", i)
//		counter++
//	}
//	lock.Unlock()
//}
