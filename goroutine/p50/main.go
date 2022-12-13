package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

// 在一些请求-响应场景中，由于各种原因，一个请求可能需要很长时间才能响应，有时甚至永远不会响应。对于这种情况，我们应该使用超时解决方案向客户端返回错误信息。这样的超时解决方案可以用select机制来实现

func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	// May need a long time to get the response.
	go doRequest(c)

	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout * time.Second):
		log.Print("111")
		return 0, errors.New("timeout")
	}
}

func doRequest(c chan<- int) {
	ra := rand.Intn(5) + 1
	log.Print("start")
	time.Sleep(time.Duration(ra) * time.Second)
	c <- 10068
	log.Print("end")
}
func main() {
	rand.Seed(time.Now().UnixNano())

	res, err := requestWithTimeout(2)

	time.Sleep(5 * time.Second)

	log.Printf("res = %d, err = %v", res, err)
}
