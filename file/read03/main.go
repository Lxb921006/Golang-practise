package main

import (
	"fmt"
	"io/ioutil"
)

var (
	count = 0
)

func main() {
	path := "C:/Users/Administrator/Desktop/test/"
	Find(path)
	fmt.Println("count = ", count)
}

func Find(path string) {
	fl, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				Find(path + file.Name() + "/")
			}
			if file.Name() == "test.txt" {
				count++
			}
			fmt.Println("file=", file.Name())
		}

	}

}
