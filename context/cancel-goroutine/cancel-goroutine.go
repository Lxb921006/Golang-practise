package main

import (
	"context"
)

func main() {
	var in = make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	go func(ctx context.Context, data <-chan int) {
		for {
			select {
			case <-data:
			case <-ctx.Done():
				return
			}
		}
	}(ctx, in)

	i := 0

	for {
		if i == 10 {
			break
		}
		in <- i
		i++
	}

}
