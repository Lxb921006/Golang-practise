package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	urls := []string{
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

	req := newRequest(ctx, 10)
	req.start()

	for _, v := range urls {
		res := &result{
			url: v,
		}

		task := func() *result {
			ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
			defer cancel1()

			go func() {
				http.Get(v)
				req.done <- struct{}{}
			}()

			select {
			case <-ctx1.Done():
				res.resp = "time out"
			case <-req.done:
				res.resp = "done"
			}

			return res
		}

		req.addTask(task)
	}
}

type result struct {
	url  string
	resp string
}

type task func() *result

type request struct {
	ctx      context.Context
	task     chan task
	result   chan *result
	workers  int
	finished chan struct{}
	done     chan struct{}
}

func newRequest(ctx context.Context, workers int) *request {
	return &request{
		ctx:      ctx,
		task:     make(chan task),
		result:   make(chan *result),
		finished: make(chan struct{}),
		done:     make(chan struct{}),
		workers:  workers,
	}
}

func (r *request) start() {
	for i := 0; i < r.workers; i++ {
		go r.work()
	}
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

			resp := v()
			r.result <- resp
		}
	}
}

func (r *request) addTask(t task) {
	r.task <- t
}
