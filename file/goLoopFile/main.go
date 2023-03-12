package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	fileCh   = make(chan string, 20)
	totalCh  = make(chan struct{})
	finishCh = make(chan struct{})
	recvCh   = make(chan struct{})
	total    = 0
	signle   = 0
)

func main() {

	start := time.Now()
	root := "/usr"

	go func() {
		for {
			select {
			case <-totalCh:
				total++
			default:

			}
		}
	}()

	go Loop(root, fileCh)

	for {
		select {
		case <-finishCh:
			signle--
			fmt.Println("--", signle)
			if signle == 0 {
				fmt.Printf("total = %d, time = %v\n", total, time.Since(start))
				return
			}
		case <-recvCh:
			fmt.Println("++", signle)
			signle++
		default:
		}
	}
}

func Loop(root string, f chan string) {
	defer func() { finishCh <- struct{}{}; <-f }()

	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				totalCh <- struct{}{}
			} else {
				select {
				case fileCh <- filepath.Join(root, file.Name()):
					go Loop(filepath.Join(root, file.Name()), fileCh)
					recvCh <- struct{}{}
				default:
					fmt.Printf("go here %s, gn = %d\n", filepath.Join(root, file.Name()), runtime.NumGoroutine())
				}
			}
		}
	}
}
