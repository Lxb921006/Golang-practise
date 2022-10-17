package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//文件操作
	file := "C:/Users/Administrator/Desktop/trucomaster.com.pem"
	wfile := "C:/Users/Administrator/Desktop/trucomaster.com-2.pem"
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

	for {
		str, err := reader.ReadString('\n') //读到一个换行符，就换行读+
		if err == io.EOF {                  //io.EOF 表示文件末尾
			break
		}
		if strings.Contains(str, "END") || strings.Contains(str, "BEGIN") {
			writer.WriteString(str)
			continue
		}
		writer.WriteString(str)
		fmt.Print(str)
	}
}
