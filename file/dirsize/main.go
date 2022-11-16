package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

type Pool struct {
	Work chan string
	Sem  chan bool //有缓冲的chan，限制goroutine number
	Exit chan bool
}

func (p *Pool) Worker(path string, size int, exit bool) {
	defer func() { <-p.Sem }()
	fl, err := os.ReadDir(path)
	if err == nil {
		for _, file := range fl {
			if file.IsDir() {
				p.Task(path+file.Name()+"/", size, false)
			} else {
				fz, _ := file.Info()
				if s := fz.Size(); s/1024/1024 > int64(size) {
					fmt.Printf("时间: %v, 文件名: %s, 文件大小: %dM, 协程数: %d\n", time.Now().Format("2006-01-02 15:04:05"), path+file.Name()+"/"+fz.Name(), fz.Size()/1024/1024, runtime.NumGoroutine())
				}
			}
		}
	}
	if !exit {
		<-p.Work
	}
	if exit {
		p.Exit <- true
	}
}

func (p *Pool) Task(task string, size int, exit bool) {
	select {
	case p.Work <- task:
	case p.Sem <- true:
		go p.Worker(task, size, exit)
	case <-p.Exit:
		return
	}
}

func NewPool(size int) *Pool {
	return &Pool{
		Work: make(chan string),
		Sem:  make(chan bool, size),
	}
}

func main() {
	//类似python的argparse，对传入的参数进行规范解析

	path := ""
	size := 0

	flag.StringVar(&path, "path", "", "目录名")
	flag.IntVar(&size, "size", 0, "要查找的文件大小")

	flag.Parse()
	if flag.NFlag() != 2 {
		fmt.Println(flag.ErrHelp.Error() + ", input -h for help")
		return
	}

	start := time.Now()

	pool := NewPool(20)
	for i := 0; i < 100; i++ {
		pool.Task(path, size, true)
	}

	fmt.Printf("耗时: %v\n", time.Since(start))
}
