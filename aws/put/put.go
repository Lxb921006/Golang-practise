package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/Lxb921006/Golang-practise/aws/s3"
)

var (
	wg     sync.WaitGroup
	dir    = make(chan string)
	put    = make(chan string)
	toStop = make(chan struct{}, 1)
)

func main() {

	const sender = 10
	const receiver = 20
	wg.Add(receiver)

	dst := "/Users/liaoxuanbiao/project/golang/src/github.com/Lxb921006/Golang-practise/.git"
	config := []string{"/Users/liaoxuanbiao/Documents/aws.ini", "aws-huawen-root", "sa-east-1"}
	s3api := &s3.S3Object{
		Bucket: "db-backup-huawen",
		S3Sess: s3.NewS3Sess(config...),
	}

	for range [receiver]struct{}{} {
		go func() {
			defer wg.Done()
			for v := range put {
				err := s3api.PutObject(v, filepath.Join("truco", filepath.Base(v)))
				if err != nil {
					log.Printf("%s failed to upload aws s3, esg >>> %s", filepath.Base(v), err.Error())
				}
			}
		}()
	}

	go TargetDir(dst, true)

	go func() {
		<-toStop
		close(put)
	}()

	wg.Wait()
}

func TargetDir(root string, exit bool) {
	fd, err := os.ReadDir(root)
	if err == nil {
		for _, v := range fd {
			if v.IsDir() {
				TargetDir(filepath.Join(root, v.Name()), false)
			} else {
				put <- filepath.Join(root, v.Name())
			}
		}
	}

	if exit {
		toStop <- struct{}{}
	}
}
