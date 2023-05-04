package main

import (
	"fmt"
	"log"
	"net"
)

// var (
// 	wg sync.WaitGroup
// )

func main() {
	//tcp编程
	// c/s结构(socket tcp/ip协议):两个客户端通过中转进行通信,
	// b/s结构(http协议):就像我们xian在看到的网站,如新浪等
	//tcp/ip协议:中文译名为传输控制协议,由网络层的ip协议跟传输层的tcp协议组成,是Internet最基本的协议
	//tcp/ip四层模型:数据链路层,网络层(ip),传输层(tcp),应用层(http,ftp)
	Server()
}

func Server() {
	fmt.Println("开始监听...")
	l, err := net.Listen("tcp", ":8092")
	if err != nil {
		log.Println("listen 8092 failed, esg >>>", err)
		return
	}

	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			return
		}
	}(l)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("客户端=%v已连接\n", c.RemoteAddr().String())
		}
		for i := 0; i < 2; i++ {
			go Process(c)
		}
	}
}

func Process(con net.Conn) {
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			return
		}
	}(con)

	for {
		buf := make([]byte, 1024)
		//读取客户端发来的数据, 如果客户端一直没发消息会阻塞,会出现超时
		n, err := con.Read(buf)
		if err != nil {
			log.Printf("客户端=%v已退出\n", con.RemoteAddr().String())
			return
		}
		//显示到终端
		log.Printf("客户端=%v, 接收到的消息=%v", con.RemoteAddr().String(), string(buf[:n]))
		//回复
		_, err = con.Write([]byte("ok"))
		if err != nil {
			return
		}
	}
}
