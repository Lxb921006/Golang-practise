package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CountCdnLog struct {
	Succeed  int
	Failed   int
	Total    int
	FileList []string
}

func (c *CountCdnLog) GetPathFile(path string) error {
	d1func := func(s string, f fs.FileInfo, err error) error {
		if f == nil {
			return err
		}

		s = strings.ReplaceAll(s, "\\", "/")
		if s == path {
			return nil
		}

		if !f.IsDir() {
			c.FileList = append(c.FileList, s)
		}
		return nil
	}
	return filepath.Walk(path, d1func)
}

func (c *CountCdnLog) CountMain(path string) error {
	err := c.GetPathFile(path)
	if err != nil {
		return err
	}
	for _, v := range c.FileList {
		f1, _ := os.OpenFile(v, os.O_RDONLY, 0777)
		defer f1.Close()
		f2 := bufio.NewReader(f1)
		for {
			f3, _, err3 := f2.ReadLine()
			if err3 == io.EOF {
				break
			}

			f5 := strings.Split(string(f3), " ")
			url := "https://hotupdate-static.burstedgold.com/hw/Biubiufishing.apk"
			matchUrl := strings.Contains(f5[6], url)
			method := strings.ReplaceAll(f5[5], "\"", "")
			status, _ := strconv.Atoi(f5[8])

			if matchUrl && method == "GET" {
				c.Total++
				if status == 200 {
					c.Succeed++
				}
				if status >= 400 && status <= 600 {
					c.Failed++
				}
			}
		}
	}
	return nil
}

func main() {
	//统计cdn log
	path := "C:/Users/Administrator/Desktop/20220415"
	fmt.Println("--------------------------")
	c := CountCdnLog{}
	err := c.CountMain(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("total=", (&c).Total)
	fmt.Println("succeed=", c.Succeed)
	fmt.Println("failed=", c.Failed)
}
