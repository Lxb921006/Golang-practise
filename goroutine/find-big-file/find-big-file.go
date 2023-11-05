package main

import (
	"fmt"
	"io/fs"
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
	root := "C:\\Windows"

	go func() {
		for range totalCh {
			total++
		}
	}()

	Loop(root, limitCh, true)

	wg.Wait()

	close(totalCh) // 关闭totalCh通道，通知统计goroutine退出

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
					go func(file fs.DirEntry) {
						defer func() {
							<-limit
							wg.Done()
						}()
						Loop(filepath.Join(root, file.Name()), limit, false)
					}(file)
				default:
					Loop(filepath.Join(root, file.Name()), limit, true)
				}
			}
		}
	}

	if !f {
		wg.Done()
	}
}
