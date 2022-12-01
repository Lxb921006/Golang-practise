package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// 搜索某个目录下的指定文件有多少个
var (
	Match     = 0
	Total     = 0
	MatchChan = make(chan bool)
	TotalChan = make(chan bool)
	// MaxWorkersChan = make(chan string, 20)
	wg sync.WaitGroup
)

func main() {
	start := time.Now()
	path := "C:/Windows/"
	// path := "C:/Users/Administrator/Desktop/test/2"
	filename := "test.txt"
	go func() {
		for {
			select {
			case <-MatchChan:
				Match++
			case <-TotalChan:
				Total++
			}
		}
	}()
	FindFile(path, filename, true)
	wg.Wait()
	close(MatchChan)
	// close(MaxWorkersChan)
	close(TotalChan)
	fmt.Printf("total = %d, count = %d, cost = %v\n", Total, Match, time.Since(start))
}

func FindFile(path, filename string, s bool) {

	// if !s {
	// 	<-MaxWorkersChan
	// }

	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.Name() == filename {
				MatchChan <- true
			}
			if file.IsDir() {
				// MaxWorkersChan <- path + file.Name() + "/" //这里超过容量，就会报deadlock!
				wg.Add(1)
				// time.Sleep(time.Second)
				go FindFile(path+file.Name()+"/", filename, false)
			} else {
				TotalChan <- true
			}
		}
	}
	if !s {
		wg.Done()
	}
}
