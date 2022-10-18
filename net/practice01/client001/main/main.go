package main

import (
	"fmt"

	UserProcesser "github.com/Lxb921006/Golang-practise/net/practice01/client001/process"
	UserModel "github.com/Lxb921006/Golang-practise/net/practice01/model"
)

var u UserModel.UserInfor

func main() {
	menu()
}

func menu() {
	ml := []string{"登录聊天系统", "注册用户", "退出系统"}
	id := 0
	out := true
	for out {
		fmt.Println("-----------欢迎登录聊天系统-----------")
		for i, v := range ml {
			fmt.Println("\t", i, v)
		}
		fmt.Printf("\t请输入%d-%d:", 0, len(ml)-1)
		fmt.Scanln(&id)
		if id < 0 || id > len(ml)-1 {
			fmt.Println()
			fmt.Printf("\t输入错误,请输入%d-%d\n", 0, len(ml)-1)
			continue
		}
		switch ml[id] {
		case "登录聊天系统":
			fmt.Println("请输入账号:")
			fmt.Scanln(&u.UserId)
			fmt.Println("请输入密码:")
			fmt.Scanln(&u.Passwd)
			up := &UserProcesser.UserProcessor{}
			res := up.Login(u.UserId, u.Passwd)
			if res != nil {
				fmt.Println("登录失败信息 err = ", res)
			} else {
				UserProcesser.LoggedMenu()
			}
		case "注册用户":
			fmt.Println("请输入账号:")
			fmt.Scanln(&u.UserId)
			fmt.Println("请输入密码:")
			fmt.Scanln(&u.Passwd)
			up := &UserProcesser.UserProcessor{}
			res := up.Register(u.UserId, u.Passwd)
			if res != nil {
				fmt.Println("注册失败信息 err = ", res)
			} else {
				fmt.Println("注册成功,请重新登录")
			}
		case "退出系统":
			fmt.Println("\tbye")
			out = false
		}
	}
}
