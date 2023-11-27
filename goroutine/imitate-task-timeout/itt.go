package main

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	resp string
	id   int
}

type pool struct {
	workers int
	wg      *sync.WaitGroup
	lock    *sync.Mutex
	once    *sync.Once
	taskCh  chan func() *result
	ctx     context.Context
	done    chan *result
}

func newPool(w int, ctx context.Context) *pool {
	return &pool{
		workers: w,
		taskCh:  make(chan func() *result),
		done:    make(chan *result),
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

	//go func() {
	//	for {
	//		select {
	//		case <-p.ctx.Done():
	//			return
	//		case <-p.done:
	//		case <-time.After(time.Second * 2):
	//			fmt.Println("time out")
	//		default:
	//			fmt.Println("gn >>> ", runtime.NumGoroutine())
	//		}
	//	}
	//}()

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
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 2)

			go func() {
				resp := v()
				p.done <- resp
			}()

			for {
				select {
				case <-ctx.Done():
					cancel()
					//resp.resp = "time out"
				case resp := <-p.done:
					resp.resp = "done"
				}
			}
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := newPool(10, ctx)
	job := p.start()

	go func() {
		for i := 0; i < 50; i++ {
			resp := &result{
				id: i,
			}

			p.taskCh <- func() *result {
				//模拟超时
				time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
				return resp
			}
		}
		p.stop()
	}()

	job.Wait()
}
