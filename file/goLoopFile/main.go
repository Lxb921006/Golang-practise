package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	fileCh   = make(chan string)
	totalCh  = make(chan struct{})
	total    = 0
	signle   = 0
	finishCh = make(chan struct{})
	recvCh   = make(chan struct{})
)

func main() {

	start := time.Now()
	root := "C:/Windows"

	for range [20]struct{}{} {
		go func() {
			for v := range fileCh {
				recvCh <- struct{}{}
				Loop(v)
			}
		}()
	}

	go func() {
		for {
			select {
			case <-totalCh:
				total++
			default:
			}
		}
	}()

	go Loop(root)

	for {
		select {
		case <-finishCh:
			signle--
			if signle == 0 {
				fmt.Printf("total = %d, time = %v/n", total, time.Since(start))
				return
			}
		case <-recvCh:
			signle++
		default:
		}
	}

}

func Loop(root string) {
	defer func() { finishCh <- struct{}{} }()
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				totalCh <- struct{}{}
			} else {
				fileCh <- filepath.Join(root, file.Name())
			}
		}
	}
}
