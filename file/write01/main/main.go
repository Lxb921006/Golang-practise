package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//写入文件
	fn := "C:/Users/Administrator/Desktop/trucomaster.com.pem"
	// f, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// for i := 0; i < 5; i++ {
	// 	_, err2 := f.WriteString("asdadasd\n")
	// 	if err2 != nil {
	// 		log.Fatal(err2)

	// 	}
	// }
	// defer f.Close()

	//写入2,带缓存的*Writer,调用WriteString时，内容是先写到缓存,还需要调用Flush方法，将缓存内容写入到文件,否则文件为空
	f11, err11 := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0777)
	if err11 != nil {
		log.Fatal(err11)
		return
	}
	defer f11.Close()
	w := bufio.NewWriter(f11)
	for i := 0; i <= 5; i++ {
		w.WriteString("aaaaabbb\r\n") //如果记事本打开没换行就再加一个\r，如：aaaa\r\n
	}
	w.Flush()

	//覆盖重写
	f2, err3 := os.OpenFile(fn, os.O_TRUNC|os.O_WRONLY, 0777)
	if err3 != nil {
		log.Fatal(err3)
		return
	}
	defer f2.Close()
	for i := 0; i < 9; i++ {
		_, err3 := f2.WriteString("22222222222\n")
		if err3 != nil {
			log.Fatal(err3)
		}
	}

	//追加
	f3, err4 := os.OpenFile(fn, os.O_APPEND, 0777)
	if err4 != nil {
		log.Fatal(err4)
	}
	defer f3.Close()
	f3.WriteString("bbbbbbbbbb\n")

	//读取文件
	f4, err5 := os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0777)
	if err5 != nil {
		log.Fatal(err5)
		return
	}
	defer f4.Close()

	f5 := bufio.NewReader(f4)
	for {
		str, err6 := f5.ReadString('\n')
		if err6 == io.EOF {
			break
		}
		fmt.Print(str)
	}

	w2 := bufio.NewWriter(f4)
	w2.WriteString("yyyyyyyyy\n")
	w2.Flush()

}
