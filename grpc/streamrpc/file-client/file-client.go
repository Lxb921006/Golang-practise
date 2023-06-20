package main

import (
	"context"
	pb "github.com/Lxb921006/Golang-practise/grpc/streamrpc/streamrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	files := []string{"E:\\googledownload\\haozip_v6.4.0.11152_compliant.exe",
		"E:\\googledownload\\python-3.9.10-amd64.exe",
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			if err := Send(file); err != nil {
				log.Println("send err >>> ", err)
			}
		}(file)
	}
	wg.Wait()
}

func Send(file string) (err error) {
	defer wg.Done()

	conn, err := grpc.Dial(":12306", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("err111 >>> ", err)
		return
	}

	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	stream, err := c.MyMethod(context.Background())

	if err != nil {
		log.Println("err222 >>> ", err)
		return
	}

	buffer := make([]byte, 8092)

	f, err := os.Open(file)
	if err != nil {
		log.Println("err333 >>> ", err)
		return
	}

	defer f.Close()

	for {
		b, err := f.Read(buffer)
		if err == io.EOF {
			log.Println("read finished111 >>>", err)
			break
		}

		if b == 0 {
			log.Println("read finished222 >>>", err)
			break
		}

		if err = stream.Send(&pb.MyMessage{Msg: buffer[:b], Name: filepath.Base(file)}); err != nil {
			log.Println("err444 >>> ", err)
			return err
		}
	}

	stream.CloseSend()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("rec ok >>> ", err)
			break
		}

		if err != nil {
			log.Println("err555 >>> ", err)
			return err
		}

		log.Println("file md5 >>> ", resp.GetName())
	}

	return
}
