package main

import (
	"testing"
)

func TestWr(t *testing.T) {
	var flag *bool
	f1 := false
	flag = &f1

	c := make(chan int)
	exit := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		go Write(c, exit, i)
	}

	for i := 0; i < 1000; i++ {
		go Read(c, exit, flag)
	}

	for {
		if *flag {
			close(c)
			close(exit)
			break
		}
	}

	t.Logf("执行完毕")
}
