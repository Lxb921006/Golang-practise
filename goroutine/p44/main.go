package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	workers   = make(chan int, 20)
	totalChan = make(chan int)
	total     = 0
)

func getwork(path string) {
	// log.Print("gn11 = ", runtime.NumGoroutine())
	defer wg.Done()
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			// log.Print("file = ", file.Name())
			if file.Name() == "ccc.log" {
				totalChan <- 1
			}
		}
	}
	<-workers
}

func sendwork(path string, finished bool) {
	log.Print("gn22 = ", runtime.NumGoroutine())
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				wg.Add(1)
				workers <- 1
				go sendwork(path, false)
				// sendwork(filepath.Join(path, file.Name()), false)
			} else {
				totalChan <- 1
			}
		}
	}

	if !finished {
		wg.Done()
		<-workers
	}
}

func main() {
	start := time.Now()
	path := "C:/Windows/"
	// path := "C:/Users/Administrator/Desktop/test/"
	go func() {
		for {
			select {
			case <-totalChan:
				total++
			default:
				// log.Print("total = ", total)
			}
		}
	}()

	sendwork(path, true)
	wg.Wait()
	log.Printf("total = %d, cost time = %v", total, time.Since(start))
}
