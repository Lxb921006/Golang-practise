package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://docs.aws.amazon.com/zh_cn/inspector/latest/user/scanning-ec2.html#deep-inspection",
		"https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonSSMManagedInstanceCore.html",
		"https://github.com/Lxb921006/chat",
		"https://stackoverflow.com/",
		"https://docs.aws.amazon.com/systems-manager/latest/userguide/agent-install-rhel.html",
		"https://www.baidu.com/",
		"https://www.qq.com/",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var workers = 10
	var countFinishedWorker int

	req := newRequest(ctx, workers)
	req.wg.Add(1)

	req.start()

	go func() {
		for {
			select {
			case <-req.ctx.Done():
				return
			case <-req.finished:
				countFinishedWorker++
				if countFinishedWorker == workers {
					close(req.result)
					return
				}

			}
		}
	}()

	go func() {
		defer req.wg.Done()
		for {
			select {
			case <-req.ctx.Done():
				return
			case v, ok := <-req.result:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}()

	go func() {
		for _, v := range urls {
			req.task <- v
		}
		req.stop()
	}()

	req.wg.Wait()
}

type result struct {
	url  string
	resp string
}

//type task func() *result

type request struct {
	ctx      context.Context
	task     chan string
	result   chan *result
	workers  int
	finished chan struct{}
	wg       *sync.WaitGroup
	lock     *sync.Mutex
}

func newRequest(ctx context.Context, workers int) *request {
	return &request{
		ctx:      ctx,
		task:     make(chan string),
		result:   make(chan *result),
		finished: make(chan struct{}),
		wg:       new(sync.WaitGroup),
		lock:     new(sync.Mutex),
		workers:  workers,
	}
}

func (r *request) start() {
	for i := 0; i < r.workers; i++ {
		go r.work()
	}
}

func (r *request) run(u string) *result {
	r.lock.Lock()
	defer r.lock.Unlock()

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	var req = new(result)
	var done = make(chan *http.Response)

	go func() {
		resp, _ := http.Get(u)
		done <- resp
	}()

	select {
	case <-ctx1.Done():
		req.url = u
		req.resp = "time out"
	case v := <-done:
		req.url = v.Request.URL.String()
		req.resp = v.Request.Host
	}

	return req
}

func (r *request) stop() {
	close(r.task)
}

func (r *request) work() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case v, ok := <-r.task:
			if !ok {
				r.finished <- struct{}{}
				return
			}

			resp := r.run(v)
			r.result <- resp
		}
	}
}
