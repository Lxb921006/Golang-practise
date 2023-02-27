package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	limit = make(chan struct{}, 10)
)

func main() {
	dir := "/Applications/"
	count := 0
	do := func(root string) {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {

				count++
			}
			return nil
		})
	}

	go Loop(dir)
	do(dir)
	time.Sleep(time.Second * 40)
	fmt.Println("file numbers = ", count)
}

func Loop(root string) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, f := range fd {
			if f.IsDir() {
				limit <- struct{}{}
				go Loop(filepath.Join(root, f.Name()))
			} else {
				log.Print(f.Name())
			}
		}
	}
	<-limit
}
