package model

import (
	"lxb-learn/net/practice01/model"
	"net"
)

type CurrentUser struct {
	model.UserInfor
	Conn net.Conn //客户端自己也要维护一个Conn
}
