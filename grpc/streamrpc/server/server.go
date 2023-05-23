package main

import (
	"fmt"
	pb "github.com/Lxb921006/Golang-practise/grpc/streamrpc/streamrpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedStreamRpcServiceServer
}

func (s *server) SayHelloWorld(req *pb.StreamRequest, stream pb.StreamRpcService_SayHelloWorldServer) (err error) {
	for range [10]struct{}{} {
		if err = stream.Send(&pb.StreamReply{
			Message: "hello rpc",
		}); err != nil {

		}
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12306))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStreamRpcServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
