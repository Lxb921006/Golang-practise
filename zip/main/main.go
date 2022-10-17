package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	dir := "C:/Users/Administrator/Desktop/111/test"
	zipName := "C:/Users/Administrator/Desktop/111/test.zip"
	unZipName := "C:/Users/Administrator/Desktop/111/upload.zip"
	unZipFile(unZipName)
	zipFile(dir, zipName)
}

//解压
func unZipFile(unZipName string) {
	f1, err1 := zip.OpenReader(unZipName)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	defer f1.Close()

	curPath := path.Dir(unZipName)

	for _, v := range f1.File {
		if v.FileInfo().IsDir() {
			if v.FileInfo().IsDir() {
				fullPath := filepath.Join(curPath, v.Name)
				err4 := os.MkdirAll(fullPath, 0777)
				if err4 != nil {
					fmt.Println(err4)
					return
				}
			}
		} else {
			fullPathFile := filepath.Join(curPath, v.Name)
			f3, err6 := v.Open()
			f4, err7 := os.OpenFile(fullPathFile, os.O_CREATE|os.O_WRONLY, 0777)
			if err6 != nil {
				fmt.Println(err6)
				return
			}
			if err7 != nil {
				fmt.Println(err7)
				return
			}
			defer f3.Close()

			f5 := bufio.NewWriter(f4)
			_, err8 := io.Copy(f5, f3)
			if err8 != nil {
				fmt.Println(err8)
				return
			}
		}
	}
}

//压缩
func zipFile(dir, zipName string) {
	file := []string{}
	topPath := path.Dir(dir) //压缩的是test目录，要去掉test前面的路径

	//存在就先删除
	_, err := os.Stat(zipName)
	if err == nil {
		os.RemoveAll(zipName)
	}

	//创建zip文件
	z1, err1 := os.Create(zipName)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}

	//打开zip文件
	z2 := zip.NewWriter(z1)
	defer z2.Close()

	//遍历目录
	err2 := GetPathFile(dir, &file)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}

	//写入zip文件
	for _, v := range file {
		f1, _ := os.Stat(v)
		header, _ := zip.FileInfoHeader(f1)
		header.Name = strings.TrimPrefix(v, topPath)

		if f1.IsDir() {
			header.Name += `/`
		}

		header.Method = zip.Deflate

		w2, _ := z2.CreateHeader(header)
		if !f1.IsDir() {
			f2, _ := os.Open(v)
			defer f2.Close()
			io.Copy(w2, f2)
		}
	}

}

func GetPathFile(old string, file *[]string) error {
	d1func := func(s string, f fs.FileInfo, err error) error {
		if f == nil {
			return err
		}
		s = strings.ReplaceAll(s, "\\", "/")
		if s != old {
			*file = append(*file, s)
		}
		return nil
	}
	return filepath.Walk(old, d1func)
}
