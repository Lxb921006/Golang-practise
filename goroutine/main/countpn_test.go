package main

import "testing"

func TestFindPn(t *testing.T) {
	res := FindPn(80000)
	if !res {
		t.Fatalf("error:%v", res)
	}
	t.Logf("ok")
}
