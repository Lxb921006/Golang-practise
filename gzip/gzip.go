package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "D:\\project\\gin\\src\\github.com\\Lxb921006\\chatai\\chat"
	gzName := filepath.Base(dir) + ".gz"
	gzPath := filepath.Join("C:\\Users\\Administrator\\Desktop", gzName)

	fs, err := os.Create(gzPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer fs.Close()

	gz := gzip.NewWriter(fs)

	defer gz.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// 打开要压缩的文件
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 复制文件内容到 tar.Writer
		_, err = io.Copy(gz, file)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
