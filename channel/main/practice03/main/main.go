package main

import (
	"fmt"
	"sync"
)

type mapChan = chan map[int]int

func main() {
	//练习1
	num := 1000
	numChan := make(chan int, num)
	resChan := make(mapChan, num)
	var wg sync.WaitGroup
	wg.Add(1)
	go AddData(numChan)
	for i := 1; i <= 8; i++ {
		go func(c chan int, r mapChan) {
			for {
				v, ok := <-c
				if !ok {
					break
				}
				r1 := Sum(v)
				r <- map[int]int{v: r1}
			}
		}(numChan, resChan)
	}
	go func(r mapChan) {
		defer wg.Done()
		for {
			if len(r) == num {
				close(r)
				for {
					v, ok := <-r
					if !ok {
						break
					}
					for i, v2 := range v {
						fmt.Printf("res[%d]=%d\n", i, v2)
					}
				}
				break
			}
		}
	}(resChan)
	wg.Wait()
}

func AddData(c chan int) {
	for i := 1; i <= cap(c); i++ {
		c <- i
	}
	close(c)
}

func Sum(n int) int {
	if n <= 1 {
		return 1
	}
	return n + Sum(n-1)
}
