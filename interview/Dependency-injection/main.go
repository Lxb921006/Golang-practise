package main

import (
	"fmt"
	"reflect"
)

// 通过依赖注入，我们可以有效结合现有组件，并将项目拆分为多个可替换、灵活、可测试的部分，这些部分之间构成易于维护和扩展的独立体系体系结构。

type Message struct {
	text string
}

// 定义MessageService接口

type MessageService interface {
	Send(message Message)
}

// 实现MessageService接口

type EmailService struct{}

func (e *EmailService) Send(message Message) {
	fmt.Printf("Sending email to %s\n", message.text)
}

// 定义Greeter接口

type Greeter struct {
	message Message
	service MessageService
}

func (g *Greeter) Greet() {
	g.service.Send(g.message)
}

// 实现NewGreeter函数，用于根据参数创建Greeter实例

func NewGreeter(message Message, service MessageService) (*Greeter, error) {
	// 验证service是否支持MessageService接口
	if reflect.TypeOf(service).Kind() != reflect.TypeOf(&EmailService{}).Kind() {
		return &Greeter{}, fmt.Errorf("service unsupported")
	}

	return &Greeter{message, service}, nil
}

func main() {
	message := Message{text: "Hey, Developer!"}

	//使用Email服务
	emailService := &EmailService{}
	greeter, err := NewGreeter(message, emailService)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	greeter.Greet()

	//使用SMTP服务...
}
