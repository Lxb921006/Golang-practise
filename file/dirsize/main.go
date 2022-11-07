package main

import (
	"fmt"
	"os"
)

func main() {
	path := "C:"
	DirSize(path)
}

func DirSize(path string) {
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				fz, _ := file.Info()
				fmt.Printf("文件名: %s, 文件大小: %d\n", fz.Name(), fz.Size())
				DirSize(path + file.Name() + "/")
			}
		}
	}
}
