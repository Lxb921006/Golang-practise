package main

import (
	"fmt"
	"os"
)

func main() {
	var chunks = make([][]byte, 5)
	var chunk []byte
	s := "aabb cc ddv cc dd cc dd 234 asda"
	b := []byte(s)

	fmt.Println("len(b) = ", len(b))

	for len(b) >= 5 {
		chunk, b = b[:5], b[5:]
		chunks = append(chunks, chunk)
		fmt.Println("len = ", len(b))
	}

	fmt.Println(chunks)

	if len(b) > 0 {
		chunks = append(chunks, b[:])
	}

	fmt.Println(chunks)

	file := "C:/Users/Administrator/Desktop/aa.txt"

	for _, chunk := range chunks {
		fmt.Println(string(chunk))
		os.WriteFile(file, chunk, 0777)
	}

}
