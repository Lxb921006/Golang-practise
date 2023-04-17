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
	iniFile    = flag.String("ini", "", "ini file path")
	section    = flag.String("section", "", "ini section")
	region     = flag.String("region", "", "aws region")
	putSrcPath = flag.String("src", "", "upload file")
)

func main() {

	flag.Parse()
	if flag.NFlag() != 4 {
		log.Fatalln(flag.ErrHelp.Error())
	}

	start := time.Now()
	const recv = 40

	root := *putSrcPath

	config := []string{*iniFile, *section, *region}
	s3api := &s3.S3Object{
		Bucket: "db-backup-huawen",
		S3Sess: s3.NewS3Sess(config...),
	}

	for range [recv]struct{}{} {
		go func() {
			for file := range WorkChan {
				fmt.Println(file)
				err := s3api.PutObject(file, "truco/"+filepath.Base(file))
				if err == nil {
					log.Printf("%s succeed to upload aws s3", filepath.Base(file))
				} else {
					log.Printf("%s failed to upload aws s3, esg >>> %s", filepath.Base(file), err.Error())
				}
			}
		}()
	}

	LoopDir(root, limitChan, true)

	wg.Wait()

	fmt.Println("loop dir finished, will close WorkChan")

	close(WorkChan)

	fmt.Printf("time = %v\n", time.Since(start))
}

func LoopDir(root string, limit chan struct{}, finished bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, file := range fd {
			if strings.Contains(filepath.Join(root, file.Name()), "2023") {
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
