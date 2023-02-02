package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano() + 1) // +1 'cause Playground's time is fixed
	fmt.Printf("doAllWork: %v\n", doAllWork())
}

func doAllWork() error {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel() // Make sure it's called to release resources even if no errors

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				// Check if any error occurred in any other gorouties:
				select {
				case <-ctx.Done():
					log.Print("timeout")
					return // Error somewhere, terminate
				default: // Default is must to avoid blocking
				}
				result, err := work(j)
				if err != nil {
					fmt.Printf("Worker #%d during %d, error: %v\n", i, j, err)
					cancel()
					return
				}
				fmt.Printf("Worker #%d finished %d, result: %d.\n", i, j, result)
			}
		}(i)
	}
	wg.Wait()

	return ctx.Err()
}

func work(j int) (int, error) {
	time.Sleep(time.Second * 30)
	return j * j, nil
}
