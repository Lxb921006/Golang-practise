package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type CopyPath struct {
	OldFileList []string
	OldPathList []string
}

func (c *CopyPath) CreateDir(newDir string) error {
	_, err1 := os.Stat(newDir)
	if os.IsNotExist(err1) {
		err2 := os.MkdirAll(newDir, 0777)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func (c *CopyPath) CopyFile(dst, src string) (w int64, err error) {
	fn11, err11 := os.OpenFile(src, os.O_RDONLY, 0777)
	if err11 != nil {
		return -1, err11
	}

	defer fn11.Close()

	fn22, err22 := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0777)
	if err22 != nil {
		return -1, err22
	}

	defer fn22.Close()

	fn111 := bufio.NewReader(fn11)
	fn222 := bufio.NewWriter(fn22)

	return io.Copy(fn222, fn111)
}

func (c *CopyPath) GetOldPathFile(old string) error {
	d1func := func(s string, f fs.FileInfo, err error) error {
		if f == nil {
			return err
		}
		s = strings.ReplaceAll(s, "\\", "/")
		if f.IsDir() {
			if s != old {
				c.OldPathList = append(c.OldPathList, s)
			}
		} else {
			c.OldFileList = append(c.OldFileList, s)
		}
		return nil
	}
	return filepath.Walk(old, d1func)
}

func (c *CopyPath) PathCopy(new, old string) error {
	if cerr := c.Check(old, new); cerr != nil {
		return cerr
	}

	if err := c.GetOldPathFile(old); err != nil {
		return err
	}

	for _, v := range c.OldPathList {
		t1 := strings.Split(v, old)
		if errc := c.CreateDir(new + t1[1]); errc != nil {
			return errc
		}
	}

	for _, v := range c.OldFileList {
		t2 := strings.Split(v, old)
		if _, errcc := c.CopyFile(new+t2[1], v); errcc != nil {
			return errcc
		}
	}

	return nil
}

func (c *CopyPath) Check(src, drc string) error {
	if src == drc {
		return errors.New("目录名重复,请检查传入的目录名")
	}
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func main() {
	fmt.Println("--------------------------------")
	d1 := "C:/Users/Administrator/Desktop/test"
	d2 := "C:/Users/Administrator/Desktop/test002"
	copy := CopyPath{}
	err := copy.PathCopy(d2, d1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("复制完成")
}
