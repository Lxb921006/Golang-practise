package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	total = 0

	root = "D:/project/gin/src/github.com/Lxb921006/Gin-bms/.git"
)

func main() {
	start := time.Now()

	Loop(root)

	fmt.Printf("total = %d, time = %v\n", total, time.Since(start))

}

func Loop(root string) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if !file.IsDir() {
				total++
			} else {
				Loop(filepath.Join(root, file.Name()))
			}
		}
	}
}
