package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	pb "github.com/Lxb921006/Golang-practise/grpc/streamrpc/streamrpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

type server struct {
	pb.UnimplementedMyServiceServer
	work chan pb.MyService_MyMethodServer
	done chan struct{}
}

func (s *server) MyMethod(stream pb.MyService_MyMethodServer) (err error) {
	log.Println("rec data")

	go func() {
		if err = s.ProcessMsg(); err != nil {
			log.Println(err)
		}
	}()

	s.work <- stream

	<-s.done

	return
}

func (s *server) ProcessMsg() (err error) {
	log.Println("process msg")
	stream := <-s.work

	var errs error
	var wf *os.File
	var file string

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("rec finished")
			break
		}

		//if err != nil {
		//	log.Println("err111 >>>", err)
		//	return
		//}

		path := "C:\\Users\\Administrator\\Desktop"
		file = filepath.Join(path, resp.GetName())
		_, err = os.Stat(file)
		if err != nil {
			wf, err = os.Create(file)
			if err != nil {
				errs = errors.New(err.Error())
			}
		}

		if errs != nil {
			log.Println("create file err ", errs)
			return errs
		}

		wf.Write(resp.Msg)

	}

	wf.Close()

	log.Println(file, " recv ok")

	m, _ := s.FileMd5(file)

	if err = stream.Send(&pb.MyMessage{Msg: []byte("md5"), Name: m}); err != nil {
		log.Println("send err ", err)
		return
	}

	s.done <- struct{}{}

	return
}

func (s *server) FileMd5(file string) (m5 string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return
	}

	m5 = hex.EncodeToString(h.Sum(nil))

	return
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12306))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{work: make(chan pb.MyService_MyMethodServer), done: make(chan struct{})})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
