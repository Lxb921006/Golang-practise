package main

import "testing"

func TestRun(t *testing.T) {
	type args struct {
		path string
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run(tt.args.path, tt.args.size)
		})
	}
}
