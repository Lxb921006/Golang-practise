package main

import (
	"fmt"
	"os"
)

//var (
//	outDirWorkChan  = make(chan string)
//	outDirLimitChan = make(chan struct{}, 10)
//	inDirWorkChan   = make(chan string)
//	inDirLimitChan  = make(chan struct{}, 20)
//)

func main() {
	root := "D:\\"
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	}
}
