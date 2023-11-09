package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	//mb
	size = 500
)

var (
	wg      sync.WaitGroup
	totalCh = make(chan struct{})
	total   = 0
)

func main() {
	start := time.Now()
	limitCh := make(chan struct{}, runtime.NumCPU())
	root := "D:\\"

	go func() {
		for {
			select {
			case file, ok := <-totalCh:
				if !ok {
					return
				}
				fmt.Println(file)
				total++
			default:
			}
		}
	}()

	Loop(root, limitCh, true)

	wg.Wait()

	close(totalCh)

	fmt.Printf("total = %d, time = %v\n", total, time.Since(start))
	var c = 0
	for c < 10 {
		c++
		fmt.Println(runtime.NumGoroutine())
		<-time.After(time.Second * 1)
	}
}

func Loop(root string, limit chan struct{}, f bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				s, err := os.Stat(filepath.Join(root, file.Name()))
				if err == nil {
					if s.Size()/1024/1024 > size {
						totalCh <- struct{}{}
					}
				}
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
