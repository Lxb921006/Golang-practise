package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	task     = make(chan string, 20)
	workdoen = make(chan bool)
)

func Worker(path string, size int, exit bool) {
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				task <- path + file.Name() + "/"
				// Worker(path+file.Name()+"/", size, false)
			} else {
				fz, _ := file.Info()
				if s := fz.Size(); s/1024/1024 >= int64(size) {
					fmt.Printf("时间: %v, 文件名: %s, 文件大小: %dM, 协程数: %d\n", time.Now().Format("2006-01-02 15:04:05"), path+file.Name()+"/"+fz.Name(), fz.Size()/1024/1024, runtime.NumGoroutine())
				}
			}
		}
	}
	if exit {
		workdoen <- true
		// fmt.Println("finished")
	}
}

func Run(path string, size int) {
	go Worker(path, size, true)

	for {
		select {
		case t := <-task:
			go Worker(t, size, false)
		case b := <-workdoen:
			fmt.Println("b = ", b)
			if b {
				return
			}
		}
	}
}

func main() {
	path := "C:/Users/Administrator/Desktop/test/"
	size := 0

	// flag.StringVar(&path, "path", "", "目录名")
	// flag.IntVar(&size, "size", 0, "要查找的文件大小")

	// flag.Parse()
	// if flag.NFlag() != 2 {
	// 	fmt.Println(flag.ErrHelp.Error() + ", input -h for help")
	// 	return
	// }

	start := time.Now()
	Run(path, size)
	// Worker(path, size, true)
	fmt.Printf("耗时: %v\n", time.Since(start))
}
