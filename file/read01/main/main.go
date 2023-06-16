package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	//文件操作
	file := "C:\\Users\\Administrator\\Desktop\\OA.rar"
	wfile := "C:\\Users\\Administrator\\Desktop\\update\\OA.rar"
	f, err := os.OpenFile(file, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	file2, err := os.OpenFile(wfile, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file2.Close()

	defer f.Close() //读取完文件要关闭文件句柄，否则有内存泄漏-占着茅坑不拉屎

	//带缓冲的Reader读取文件
	//默认缓冲区是4096字节，也就是说读取文件不会一次性将所有文件内容加载到内存而是每次加载4096字节到内存处理,适合读取大文件
	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(file2)

	io.Copy(writer, reader)

	log.Println("ok")
}
