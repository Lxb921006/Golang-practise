package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	totalCh = make(chan struct{})
	total   = 0
)

func main() {
	limitCh := make(chan struct{}, 20)
	start := time.Now()
	root := "C:/Windows"

	go func() {
		for {
			select {
			case <-totalCh:
				total++
			default:
			}
		}
	}()

	Loop(root, limitCh, true)

	wg.Wait()

	fmt.Printf("total = %d, time = %v\n", total, time.Since(start))

}

func Loop(root string, limit chan struct{}, f bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				totalCh <- struct{}{}
			} else {
				select {
				case limit <- struct{}{}:
					wg.Add(1)
					go Loop(filepath.Join(root, file.Name()), limit, false)
				default:
					Loop(filepath.Join(root, file.Name()), limit, true)
				}
			}
		}
	}

	if !f {
		wg.Done()
		<-limit
	}
}
