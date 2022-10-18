package main

//这里是处理消息类型的总控文件,根据不同的消息类型分发给其它专门的处理逻辑函数如process目录下的userProcess.go

import (
	"fmt"
	"io"
	"net"

	Common "github.com/Lxb921006/Golang-practise/net/practice01/common"
	UserMessage "github.com/Lxb921006/Golang-practise/net/practice01/model"
	UserProcess "github.com/Lxb921006/Golang-practise/net/practice01/server001/process"
)

type Processer struct {
	Conn net.Conn
}

func (p *Processer) MessagesProcess(mes *UserMessage.Message) (err error) {
	fmt.Println("MessagesProcess mes = ", mes)
	switch mes.Type {
	case UserMessage.LoginMessAgeType:
		//处理登录
		u := &UserProcess.UserProcessor{
			Conn: p.Conn,
		}
		err = u.LoginProcess(mes)
		return
	case UserMessage.RegisterType:
		//处理注册
		u := &UserProcess.UserProcessor{
			Conn: p.Conn,
		}
		err = u.RegisterProcess(mes)
		return
	case UserMessage.ChatMessageType:
		//处理群发消息
		cp := &UserProcess.ChatProcessor{}
		cp.SendMsgToAllUsers(mes)
	case UserMessage.ChatUserToUserMessageType:
		//处理一对一聊天消息

		UserProcess.Chp.SendMsgToOneOnLingUser(mes)
	case UserMessage.UserStatusChangeType:
		//修改用户状态
		u := &UserProcess.UserProcessor{
			Conn: p.Conn,
		}
		err = u.LoginOut(mes)
	case UserMessage.FileMessageType:
		fmt.Println("1111111111111")
		fp := &UserProcess.FileProcess{}
		fp.SendFileToUser(mes)
	default:
		//消息类型不存在
		fmt.Println("消息类型不存在")
	}
	return
}

func (p *Processer) MainProcesser() (err error) {
	for {
		t := &Common.TransData{
			Conn: p.Conn,
		}
		mes, err := t.RecvMessage()
		if err != nil {
			if err == io.EOF {
				//fmt.Printf("客户端=%v, 已关闭", t.Conn.RemoteAddr().String())
				return err
			} else {
				fmt.Println("recv message err = ", err)
				return err
			}
		}
		//获取客户端发来的登录信息根据消息类型分类处理
		err = p.MessagesProcess(&mes)
		if err != nil {
			return err
		}
	}
}
