package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	resp := task(ctx)
	fmt.Println(resp)

}

func task(ctx context.Context) error {
	time.Sleep(time.Second * time.Duration(14))
	if ctx != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	return nil

}
