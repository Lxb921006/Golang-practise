package main

import "testing"

func TestFindPn2(t *testing.T) {
	wr := make(chan int, 1000)
	rwr := make(chan int, 3000)
	go Write(wr)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go Read(wr, rwr)
	}
	wg.Wait()
}
