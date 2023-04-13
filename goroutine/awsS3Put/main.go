package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Lxb921006/Golang-practise/aws/s3"
)

var (
	limitChan  = make(chan struct{}, 20)
	WorkChan   = make(chan string)
	wg         sync.WaitGroup
	stopChan   = make(chan struct{}, 1)
	iniFile    = flag.String("ini", "", "ini file path")
	section    = flag.String("section", "", "ini section")
	region     = flag.String("region", "", "aws region")
	putSrcPath = flag.String("src", "", "upload file")
)

func main() {
	start := time.Now()
	flag.Parse()
	root := *putSrcPath

	config := []string{*iniFile, *section, *region}
	s3api := &s3.S3Object{
		Bucket: "db-backup-huawen",
		S3Sess: s3.NewS3Sess(config...),
	}

	go func() {
		for {
			select {
			case file := <-WorkChan:
				err := s3api.PutObject(file, "truco/"+filepath.Base(file))
				if err == nil {
					log.Printf("%s succeed to upload aws s3", filepath.Base(file))
				} else {
					log.Printf("%s failed to upload aws s3, esg >>> %s", filepath.Base(file), err.Error())
				}
			case <-stopChan:
				return
			}
		}
	}()

	LoopDir(root, limitChan, true)

	wg.Wait()

	stopChan <- struct{}{}

	fmt.Printf("time = %v\n", time.Since(start))
}

func LoopDir(root string, limit chan struct{}, finished bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if strings.HasPrefix(filepath.Join(root, file.Name()), "sbl_") {
				if file.Name() == "MGLog" || file.Name() == "LOG" {
					continue
				}
				if file.IsDir() {
					select {
					case limit <- struct{}{}:
						wg.Add(1)
						go LoopDir(filepath.Join(root, file.Name()), limit, false)
					default:
						LoopDir(filepath.Join(root, file.Name()), limit, true)
					}
				} else {
					WorkChan <- filepath.Join(root, file.Name())
				}
			}
			
		}
	}

	if !finished {
		wg.Done()
		<-limit
	}
}
