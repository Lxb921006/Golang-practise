package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.qq.com/",
		"http://www.jd.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		// Increment the WaitGroup counter.
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			r, _ := http.Get(url)
			fmt.Printf("url:%v stauts=%v\n", url, r.Status)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("ok")
}
