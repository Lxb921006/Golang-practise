package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	resp string
	id   int
}

type pool struct {
	workers  int
	wg       *sync.WaitGroup
	lock     *sync.Mutex
	once     *sync.Once
	taskCh   chan func() *result
	ctx      context.Context
	done     chan *result
	finished chan bool
}

func newPool(w int, ctx context.Context) *pool {
	return &pool{
		workers:  w,
		taskCh:   make(chan func() *result),
		done:     make(chan *result),
		ctx:      ctx,
		wg:       new(sync.WaitGroup),
		once:     new(sync.Once),
		lock:     new(sync.Mutex),
		finished: make(chan bool, w),
	}
}

func (p *pool) start() *sync.WaitGroup {
	p.wg.Add(p.workers)
	for i := 0; i < p.workers; i++ {
		go p.work()
	}

	return p.wg
}

func (p *pool) work() {
	defer p.wg.Done()
	for {
		select {
		case <-p.ctx.Done():
			return
		case v, ok := <-p.taskCh:
			if !ok {
				p.finished <- true
				return
			}

			resp := v()
			resp.resp = "done"
			p.done <- resp
		}
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

func (p *pool) addTask(task func() *result) {
	p.taskCh <- task
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newPool(10, ctx)
	job := p.start()
	var f int

	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			case <-p.finished:
				if f == 9 {
					close(p.done)
					return
				}
				f++
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-p.ctx.Done():
				return
			case result, ok := <-p.done:
				if !ok {
					return
				}
				fmt.Printf("Job #%d is %s\n", result.id, result.resp)
			}
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			resp := &result{id: i}
			task := func() *result {
				ctx1, cancel1 := context.WithTimeout(context.Background(), 2)
				defer cancel1()

				time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)

				select {
				case <-ctx1.Done():
					resp.resp = "time out"
				}

				return resp
			}
			p.addTask(task)
		}
		p.stop()
	}()

	job.Wait()
	wg.Wait()
}
