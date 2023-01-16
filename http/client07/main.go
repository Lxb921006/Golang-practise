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
		"https://github.com/x-ream/rey",
		"https://github.com/ericchiang/k8s",
		"https://github.com/kubernetes/kops",
		"https://github.com/nats-io/k8s",
		"https://github.com/dotbalo/k8s",
		"https://github.com/Thakurvaibhav/k8s",
		"https://github.com/lucassha/CKAD-resources",
		"https://github.com/Thakurvaibhav/k8s",
		"https://github.com/k8sre/k8s",
		"https://github.com/k3s-io/k3s",
		"https://github.com/Thakurvaibhav/k8s",
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
		Timeout: time.Duration(3) * time.Second,
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
		_, err := client.Get(url)
		if err != nil {
			log.Println("We could not reach:", url, err)
		} else {
			log.Println("Success reaching the website:", url)
		}

	}

}
