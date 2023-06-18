package main

import (
	"fmt"
	pb "github.com/Lxb921006/Golang-practise/grpc/streamrpc/streamrpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
)

var (
	once sync.Once
	wf   *os.File
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyMethod(stream pb.MyService_MyMethodServer) (err error) {
	log.Println("rec data")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("rec finished")
			break
		}

		once.Do(func() {
			path := "C:\\Users\\Administrator\\Desktop"
			file := filepath.Join(path, resp.GetName())
			wf, err = os.Create(file)
			if err != nil {
				log.Fatal("err3 >>> ", err)
			}
		})

		wf.Write(resp.Msg)

	}

	log.Println("write ok")

	//if err = stream.Send(&pb.MyMessage{Msg: []byte("md5"), Name: fn}); err != nil {
	//	return
	//}
	wf.Close()

	return
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12306))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
