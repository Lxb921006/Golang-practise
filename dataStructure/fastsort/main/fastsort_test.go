package main

import (
	"math/rand"
	"testing"
)

func TestFastSort(t *testing.T) {
	num := []int{}
	for i := 0; i < 800000; i++ {
		num = append(num, rand.Intn(9000000))
	}

	FastSort(num, 0, len(num)-1)
}
