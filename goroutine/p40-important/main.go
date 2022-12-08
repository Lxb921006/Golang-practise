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
	totalChan = make(chan int)
	total     = 0
)

func getwork() {
	for path := range work {
		// log.Print("gn111 = ", runtime.NumGoroutine())
		fl, err := os.ReadDir(path)
		if err == nil {
			for _, file := range fl {
				if !file.IsDir() {
					totalChan <- 1
					// log.Print(file.Name())
				}
			}
		}
	}
}

func sendwork(path string, finished bool) {
	// log.Print("gn222 = ", runtime.NumGoroutine())
	fl, err := os.ReadDir(path)
	if err == nil {
		work <- path
		for _, file := range fl {
			if file.IsDir() {
				sendwork(filepath.Join(path, file.Name()), false)
			}
		}
	}

	if finished {
		close(work)
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
				log.Print("total =", total)
			}
		}
	}()

	for i := 0; i < 20; i++ {
		go getwork()
	}

	sendwork(path, true)

	// for {
	// 	if runtime.NumGoroutine() <= 2 {
	// 		break
	// 	}
	// }

	var block chan int
	<-block

	log.Printf("total = %d, cost time = %v", total, time.Since(start))

}
