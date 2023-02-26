package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "/Users/liaoxuanbiao/project/golang/src/github.com/Lxb921006/Golang-practise/.git"
	count := 0
	do := func(root string) {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Println(info.Name())
				count++
			}
			return nil
		})
	}

	do(dir)
	fmt.Println("file numbers = ", count)
}
