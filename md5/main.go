package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file := "D:\\工作工具\\SQLServer2019-x64-CHS.iso"
	f, _ := os.Open(file)
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	md5sum := hex.EncodeToString(h.Sum(nil))
	fmt.Println(md5sum)
	//fmt.Printf("%x\n", h.Sum(nil))
}
