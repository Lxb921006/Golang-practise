package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fn := "C:/Users/Administrator/Desktop/11111.txt"
	err := PathCheck(fn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v exists\n", fn)

	f1, _ := os.Stat(fn)
	if f1.IsDir() {
		fmt.Println("目录")
	} else {
		fmt.Println("文件")
	}

	fmt.Println(os.Getwd())
	// s := regexp.MustCompile("/").Split(fn, -1)
	// getPath := strings.Join(s[:len(s)-1], "/")
	getPath, fileName := path.Split(fn) //返回文件路径，文件名
	fmt.Println(getPath)
	fmt.Println(fileName)
	fmt.Println(path.Join("a", "b", "c")) //拼接路径
	fmt.Println(path.Ext(fn))             //文件扩展名
	fmt.Println("----------------------------------------")
	//打印出目录下的所有文件
	dir := "C:/Users/Administrator/Desktop/test"
	// res := PathWalk(dir)
	// if res != nil {
	// 	return
	// }
	path, file := []string{}, []string{}

	err1 := PathWalk(dir, &path, &file)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(path)
	fmt.Println(file)
}

func PathCheck(path string) error {
	_, err := os.Stat(path) //直接判断这个err是否为nil也可以
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return errors.New(err.Error())
	}
	return errors.New(err.Error())
}

func PathWalk(path string, dir, file *[]string) error { //遍历目录
	d1func := func(s string, f fs.FileInfo, err error) error {
		if f == nil {
			return err
		}
		s = strings.ReplaceAll(s, "\\", "/")
		if f.IsDir() {
			if s != path {
				*dir = append(*dir, s)
			}
		} else {
			*file = append(*file, s)
		}
		return nil
	}
	return filepath.Walk(path, d1func)
}
