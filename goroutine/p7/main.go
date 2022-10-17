package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var (
	match = 0
	total = 0
)

func main() {

	start := time.Now()
	// path := "C:/Users/Administrator/Desktop/test/"
	path := "C:/Users/"
	FindFile(path, "test.txt")
	fmt.Printf("total = %d, count file = %d,cost time = %v\n", total, match, time.Since(start))
}

func FindFile(path, filename string) {
	fl, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				FindFile(path+file.Name()+"/", filename)
			} else {
				if file.Name() == filename {
					match++
				}
				// total++
			}

		}
	}
}
