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
)

func main() {
	conn, err := grpc.Dial(":12306", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	stream, err := c.MyMethod(context.Background())
	if err != nil {
		log.Fatal("err1 >>> ", err)
	}

	file := "E:\\googledownload\\python-3.9.10-amd64.exe"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("err2 >>> ", err)
	}

	defer f.Close()

	buffer := make([]byte, 8092)

	defer stream.CloseSend()

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

		if err := stream.Send(&pb.MyMessage{Msg: buffer[:b], Name: filepath.Base(file)}); err != nil {
			log.Fatal("err3 >>> ", err)
		}
	}

}
