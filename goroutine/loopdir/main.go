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
	limitCh := make(chan struct{}, 20)
	root := "C:\\Windows"

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

	Loop(root, limitCh, true)

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
					// 在limit阻塞期间, 让Loop自己也可以继续遍历出文件
					Loop(filepath.Join(root, file.Name()), limit, true)
				}
			}
		}
	}

	// 遍历完目录要给让limit给出位置给新的goroutine用,并标记当前goroutine已完成
	if !f {
		<-limit
		wg.Done()
	}
}
