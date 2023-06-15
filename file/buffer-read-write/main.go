package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// 一般是一次读取4K或者8K的数据到缓冲区
func main() {

	start := time.Now()

	var rb = make([]byte, 8092)
	file := "E:\\googledownload\\goland-2023.1.1.exe"
	path := "C:\\Users\\Administrator\\Desktop"

	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	output := filepath.Join(path, "test.exe")

	fn, err := os.Create(output)
	if err != nil {
		return
	}

	for {
		n, err := f.Read(rb)
		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		_, err = fn.Write(rb[:n])
		if err != nil {
			return
		}
	}

	fmt.Println("done, cost time = ", time.Since(start))
}
