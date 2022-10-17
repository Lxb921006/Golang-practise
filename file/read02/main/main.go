package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f := "C:/Users/Administrator/Desktop/事项.txt"
	fmt.Println("----------------------------------")
	//这里没有open,所以不需要close,适合读取小文件

	s, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", s)
	fmt.Println("----------------------------------")
	s2, err2 := os.ReadFile(f)
	if err != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%v\n", len(s2))

	b, _ := json.Marshal(&s2)

	d := uint32(len(b))

	buf := make([]byte, 1024)
	binary.LittleEndian.PutUint32(buf[:4], d)

	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", buf[:4])
}
