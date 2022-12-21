package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {

	filename := "/Users/liaoxuanbiao/Downloads/log/20221216/2022121618.log.gz"

	gzipfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()
}
