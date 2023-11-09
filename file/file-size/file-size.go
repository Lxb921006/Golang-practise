package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	root := "D:\\"

	fl, _ := os.ReadDir(root)
	for _, file := range fl {
		if !file.IsDir() {
			s, err := os.Stat(filepath.Join(root, file.Name()))
			if err == nil {
				if s.Size()/1024/1024 > 50 {
					fmt.Printf("file: %s, size: %dmb\n", s.Name(), s.Size()/1024/1024)
				}
			}
		}
	}

	fmt.Println("cpu num >>> ", runtime.NumCPU())
}
