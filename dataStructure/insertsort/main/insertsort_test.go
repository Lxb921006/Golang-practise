package main

import (
	"math/rand"
	"testing"
)

func TestFastSort(t *testing.T) {
	num := [80000]int{}
	for i := 0; i < 80000; i++ {
		num[i] = rand.Intn(9000000)
	}

	insertsort(&num)
}
