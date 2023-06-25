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
	files := []string{"D:\\工作工具\\TortoiseSVN64.msi",
		"D:\\工作工具\\天锐绿盾终端.exe",
		"D:\\工作工具\\SQLServer2019-x64-CHS.iso",
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
		return
	}

	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	stream, err := c.MyMethod(context.Background())

	if err != nil {
		return
	}

	buffer := make([]byte, 8092)

	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	for {
		b, err := f.Read(buffer)
		if err == io.EOF {
			break
		}

		if b == 0 {
			break
		}

		if err = stream.Send(&pb.MyMessage{Msg: buffer[:b], Name: filepath.Base(file)}); err != nil {
			return err
		}
	}

	stream.CloseSend()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Println("file md5 >>> ", resp.GetName())
	}

	return
}
