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
	limitCh = make(chan struct{}, 20)
	totalCh = make(chan struct{})
	total   = 0
)

func main() {

	start := time.Now()
	root := "C:/Windows"

	go func() {
		for {
			select {
			case <-totalCh:
				total++
			default:
				// fmt.Println("gn = ", runtime.NumGoroutine())
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
				// wg.Add(1)
				// go Loop(filepath.Join(root, file.Name()), limitCh, false)
				select {
				case limitCh <- struct{}{}:
					wg.Add(1)
					go Loop(filepath.Join(root, file.Name()), limitCh, false)
				default:
					Loop(filepath.Join(root, file.Name()), limitCh, true)
				}
			}
		}
	}

	if !f {
		wg.Done()
		<-limit
	}
}
