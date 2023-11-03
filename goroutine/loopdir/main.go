package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	limitCh := make(chan struct{}, 200)
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

	Loop(root, limitCh, true)

	wg.Wait()

	//防止内存泄露
	close(totalCh)

	fmt.Printf("total = %d, time = %v\n", total, time.Since(start))
}

func Loop(root string, limit chan struct{}, f bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				fmt.Printf("time >>> %s, count >>> %d, gn >>> %d, channel_cap >>> %d\n", time.Now().Format("2006-01-02 15:04:05"), total, runtime.NumGoroutine(), len(limit))
				totalCh <- struct{}{}
			} else {
				select {
				case limit <- struct{}{}: // 限制goroutine创建数量
					wg.Add(1)
					go Loop(filepath.Join(root, file.Name()), limit, false)
				default:

					// 在limit阻塞期间, 让Loop自己也可以继续遍历出文件统计
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
