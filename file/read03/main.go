package main

import (
	"fmt"
	"os"
)

func main() {
	file := "C:\\Users\\Administrator\\Desktop\\update\\天锐绿盾终端1.exe"
	_, err := os.Stat(file)
	fmt.Println(err)

}
