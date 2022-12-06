package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type Seat int
type Bar chan Seat

var (
	// wg      sync.WaitGroup
	work      = make(chan string)
	workers   = make(chan int, 20)
	totalChan = make(chan int)
	total     = 0
)

func getwork() {
	for path := range work {
		fl, err := os.ReadDir(path)
		if err == nil {
			for _, file := range fl {
				if !file.IsDir() {
					totalChan <- 1
				}
			}
		}
	}
}

func sendwork(path string, finished bool) {
	fl, err := os.ReadDir(path)
	if err == nil {
		work <- filepath.Join(path)
		for _, file := range fl {
			if file.IsDir() {
				go sendwork(filepath.Join(path, file.Name()), false)
			}
		}
	}

	// if finished {
	// 	close(work)
	// }
}

func main() {
	path := "C:/Windows/"
	start := time.Now()
	// path := "C:/Users/Administrator/Desktop/test/"
	go func() {
		for {
			select {
			case <-totalChan:
				total++
			default:
			}
		}
	}()

	for i := 0; i < cap(workers); i++ {
		go getwork()
	}

	sendwork(path, true)

	for {
		// log.Print("gn = ", runtime.NumGoroutine())
		log.Printf("total = %d, cost time = %v", total, time.Since(start))
	}

	log.Printf("total = %d, cost time = %v", total, time.Since(start))

}
