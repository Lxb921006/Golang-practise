package main

import (
	"fmt"
	"lxb-learn/net/practice01/server001/model"
	"lxb-learn/net/practice01/server001/process"
	"net"
)

func init() {
	//初始化redis连接池
	InitPoolRds("101.35.143.86:6378")
	//这里是初始化usermodel的实例
	InitUserCdus()
}

func main() {
	Server()
}

//这里是初始化usermodel的实例
func InitUserCdus() {
	model.IuserCdus = model.NewPoolRds(RdPool)
}

func Server() {
	fmt.Println("开始监听...")
	l, e1 := net.Listen("tcp", "0.0.0.0:8092")
	if e1 != nil {
		fmt.Println("listen 8092 failed")
		return
	}
	defer l.Close()

	for {
		c, e2 := l.Accept()
		if e2 != nil {
			fmt.Println(e2)
		} else {
			fmt.Printf("客户端=%v已连接\n", c.RemoteAddr().String())
		}
		go Process(c)
	}
}

func Process(con net.Conn) {
	defer con.Close()
	p := &Processer{
		Conn: con,
	}
	//处理消息的总控函数
	err := p.MainProcesser()
	if err != nil {
		fmt.Println("online-1 = ", process.UserMgr)
		fmt.Printf("客户端=%v已关闭, msg= %v\n", p.Conn.RemoteAddr().String(), err)
		return
	}
	fmt.Println("online-2 = ", process.UserMgr)

}
