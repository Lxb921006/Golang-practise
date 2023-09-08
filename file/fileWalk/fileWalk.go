package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {

	dir := "D:\\project\\gin\\src\\github.com\\Lxb921006\\chatai\\chat\\.git"

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			//fileName := filepath.Join(path, info.Name())
			fmt.Println(path, info.Name())
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
