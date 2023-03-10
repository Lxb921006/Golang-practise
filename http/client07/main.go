package main

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	work = make(chan string)
)

func main() {

	links := []string{
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
		"http://161.117.87.119:10086/user/list/",
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Duration(5) * time.Second,
	}

	const MaxWorkers = 4
	wg.Add(MaxWorkers)

	for range [MaxWorkers]struct{}{} {
		go request(client)
	}

	for _, url := range links {
		work <- url
	}

	close(work)

	wg.Wait()

}

func request(client *http.Client) {
	defer wg.Done()

	for url := range work {

		log.Print("goroutine number = ", runtime.NumGoroutine())
		resp, err := client.Get(url)
		if err != nil {
			log.Println("We could not reach:", url, err, resp.StatusCode)
		} else {
			log.Println("Success reaching the website:", url, resp.StatusCode)
		}

	}

}
