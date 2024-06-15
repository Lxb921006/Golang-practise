package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var send = make(chan string)
var wg sync.WaitGroup

// var wgDir sync.WaitGroup
var worker = 2
var finish = make(chan string)

func main() {
	now := time.Now()
	//cmd := "rsync"
	cmd := "D:\\rsync\\ICW\\bin\\rsync.exe"
	//host := "10.0.0.219::nas/id-db-bak/"
	host := "192.168.3.11::web/shell/"
	//dirs := []string{
	//	//"Z:\\dbbak\\thailDB$thgroup",
	//	//"Z:\\dbbak\\TH-DB-2",
	//	"Z:\\dbbak\\id-cluster$idgroup",
	//	"Z:\\dbbak\\ID-DB2",
	//}
	dirs := []string{"C:\\Users\\Administrator\\Desktop\\111", "C:\\Users\\Administrator\\Desktop\\222"}
	dirsLen := len(dirs)
	var over = make(chan struct{}, dirsLen)
	ctx := context.Background()

	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			ctx1, cancel := context.WithTimeout(ctx, time.Second*time.Duration(7200))
			defer cancel()
			for file := range send {
				if err := exec.CommandContext(ctx1, cmd, "-av", "--ignore-errors", file, host).Run(); err != nil {
					log.Println("exec cmd err >>> ", err, file)
				} else {
					fmt.Printf("successful %s \n", file)
				}
			}
		}(ctx)
	}

	go func() {
		for v := range finish {
			loopDir(v)
			over <- struct{}{}
		}
	}()

	go func() {
		for {
			select {
			case <-over:
				if dirsLen == 1 {
					close(send)
					return
				}
				dirsLen--
			}
		}
	}()

	for _, dir := range dirs {
		finish <- dir
	}

	close(finish)
	wg.Wait()
	close(over)

	fmt.Printf("传输总共耗时 >>> %s", time.Now().Sub(now))
}

func loopDir(dir string) {
	readDir, err := os.ReadDir(dir)
	if err == nil {
		for _, v := range readDir {
			if v.IsDir() {
				loopDir(filepath.Join(dir, v.Name()))
			} else {
				if strings.HasSuffix(filepath.Join(dir, v.Name()), ".bak") && !strings.Contains(filepath.Join(dir, v.Name()), "20240614") && !strings.Contains(filepath.Join(dir, v.Name()), "test") {
					rf := replace(filepath.Join(dir, v.Name()))
					send <- rf
				}
			}
		}
	}

}

func replace(path string) string {
	file := strings.Split(path, ":")
	pf := filepath.Join("/cygdrive/z/", file[1])
	return strings.ReplaceAll(pf, "\\", "/")
}
