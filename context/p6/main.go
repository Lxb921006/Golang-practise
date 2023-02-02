package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

var (
	stop = make(chan int, 1)
)

func main() {
	do := make(chan int)

	rand.Seed(time.Now().Unix())
	s := []int{100, 200, 300, 400, 500, 600}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	work := func(ctx context.Context) {
		for v := range do {
			run(v, ctx)
		}
	}

	for range [3]struct{}{} {
		go work(ctx)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				stop <- 1
				log.Print("timeout222")
			case <-time.After(2 * time.Second):
				stop <- 1
				log.Print("timeout111")
			}
		}
	}()

	for _, v := range s {
		do <- v
	}

	time.Sleep(time.Second * 20)

}

func run(v int, ctx context.Context) {
	for {
		select {
		case <-stop:
			log.Print(v)
			return
		default:
		}
	}
	log.Print(111)
}
