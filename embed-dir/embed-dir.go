package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed dir

var data embed.FS

func main() {
	dir, _ := data.ReadDir("dir")
	for _, v := range dir {
		f, _ := v.Info()
		fmt.Println(f.Name())
	}
}
