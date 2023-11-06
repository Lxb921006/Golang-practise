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
	start := time.Now()
	limitCh := make(chan struct{}, 10)
	root := "D:\\"

	go func() {
		for {
			select {
			case _, ok := <-totalCh:
				if !ok {
					return
				}
				total++
			default:
			}
		}
	}()

	fd, _ := os.ReadDir(root)
	for _, fn := range fd {
		if fn.IsDir() {
			fmt.Println(filepath.Join(root, fn.Name()))
			Loop(filepath.Join(root, fn.Name()), limitCh, true)
		}
	}

	wg.Wait()

	close(totalCh)

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
