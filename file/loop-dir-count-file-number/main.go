package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	total = 0

	root = "C:/Users/Administrator/Desktop/update"
)

func main() {
	start := time.Now()

	Loop(root)

	log.Println("file numbers = ", total, "time = ", time.Since(start))

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
