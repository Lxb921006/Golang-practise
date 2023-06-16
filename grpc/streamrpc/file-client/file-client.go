package main

import (
	"bufio"
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

	file := "C:\\Users\\Administrator\\Desktop\\OA.rar"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("err2 >>> ", err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		str, err := reader.ReadString('\n') //读到一个换行符，就换行读+
		if err == io.EOF {                  //io.EOF 表示文件末尾
			break
		}

		if err := stream.Send(&pb.MyMessage{Msg: []byte(str), Name: filepath.Base(file)}); err != nil {
			log.Fatal("err3 >>> ", err)
		}
	}

	//for {
	//	resp, err := stream.Recv()
	//	if err == io.EOF {
	//		fmt.Println("rec ok")
	//		break
	//	}
	//
	//	if err != nil {
	//		log.Fatal("err4 >>> ", err)
	//	}
	//
	//	log.Println(string(resp.GetMsg()))
	//}

}
