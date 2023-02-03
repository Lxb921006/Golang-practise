package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Channel used to receive the result from doSomething function
	ch := make(chan string, 1)

	// Create a context with a timeout of 5 seconds
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Start the doSomething function
	for range [4]struct{}{} {
		go func() {
			doSomething(ctxTimeout, ch)
		}()
	}

	go func() {
		for {
			select {
			case <-ctxTimeout.Done():
				fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
				return
			case result := <-ch:
				fmt.Printf("Received: %s\n", result)
			}
		}

	}()

	time.Sleep(30 * time.Second)
}

func doSomething(ctx context.Context, ch chan string) {
	fmt.Println("doSomething Sleeping...")
	run()
	fmt.Println("doSomething Wake up...")
	ch <- "Did Something"
}

func run() {
	for {
		time.Sleep(time.Second)
	}
}
