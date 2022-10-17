package main

import "testing"

func TestFindPn2(t *testing.T) {
	for i := 1; i <= 80000; i++ {
		wr <- i
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// time.Sleep(time.Second)
			FindPn2(i, 80000)
			<-wr
			<-rwr
		}(i)
	}
	wg.Done()
}

func TestFindPn3(t *testing.T) {
	for i := 1; i <= 80000; i++ {
		FindPn2(i, 80000)
	}
}
