package main

import "testing"

func TestSort1(t *testing.T) {
	res := Sort1()
	t.Logf("res1=%v", res)
}

func TestSort2(t *testing.T) {
	res := Sort2()
	t.Logf("res2=%v", res)
}

func TestFindPn2(t *testing.T) {
	for i := 1; i <= 80000; i++ {
		FindPn2(i, 80000)
	}
}
