package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	//类似python的argparse，对传入的参数进行规范解析
	start := time.Now()
	path := ""
	size := 0

	flag.StringVar(&path, "path", "", "目录名")
	flag.IntVar(&size, "size", 0, "要查找的文件大小")

	flag.Parse()
	if flag.NFlag() != 2 {
		fmt.Println(flag.ErrHelp.Error() + ", input -h for help")
		return
	}
	DirSize(path, size, true)
	wg.Wait()

	fmt.Printf("耗时: %v\n", time.Since(start))
}

func DirSize(path string, size int, flag bool) {
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				wg.Add(1)
				go DirSize(path+file.Name()+"/", size, false)
			} else {
				fz, _ := file.Info()
				if s := fz.Size(); s/1024/1024 > int64(size) {
					fmt.Printf("时间: %v, 文件名: %s, 文件大小: %dM\n", time.Now().Format("2006-01-02 15:04:05"), path+file.Name()+"/"+fz.Name(), fz.Size()/1024/1024)
				}
			}
		}
	}
	if !flag {
		wg.Done()
	}
}
