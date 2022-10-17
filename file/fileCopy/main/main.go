package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//文件复制
	fn1 := "C:/Users/Administrator/Desktop/天锐绿盾终端.exe"
	fn2 := "C:/Users/Administrator/Desktop/test/天锐绿盾终端.exe"
	//复制文件
	_, err1 := CopyFile(fn2, fn1)
	if err1 != nil {
		fmt.Println(err1)
	}

}

func CopyFile(dst, src string) (w int64, err error) {
	fn11, err11 := os.OpenFile(src, os.O_RDONLY, 0777)
	if err11 != nil {
		return -1, err11
	}

	defer fn11.Close()

	fn22, err22 := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0777)
	if err22 != nil {
		return -1, err22
	}

	defer fn22.Close()

	fn111 := bufio.NewReader(fn11)
	fn222 := bufio.NewWriter(fn22)

	return io.Copy(fn222, fn111) //可以处理大文件，底层是带缓存的
}
