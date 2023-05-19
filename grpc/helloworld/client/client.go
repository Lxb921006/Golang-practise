package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Lxb921006/Golang-practise/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "lxb", "Name to sayHelloWorld")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewTestGrpcHelloWorldClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for range [10]struct{}{} {
		r, err := c.SayHelloWorld(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not send: %v", err)
		}
		log.Printf("recv: %s", r.GetMessage())
		time.Sleep(time.Second / 2)
	}

}
