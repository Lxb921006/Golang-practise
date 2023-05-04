package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"
)

func main() {
	Client()
}

func Client() {
	//tcp编程客户端
	c, err := net.Dial("tcp", "127.0.0.1:8092")
	if err != nil {
		log.Println(err)
		return
	}

	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			return
		}
	}(c)

	//发送单行数据
	reader := bufio.NewReader(os.Stdin)
	exContent := regexp.MustCompile(`exit`) //strings.Trim(content," \n\r")

	for {

		buf := make([]byte, 1024)

		//从终端获取输入的数据
		content, _ := reader.ReadString('\n')

		if exContent.MatchString(content) {
			break
		}

		//将输入的内容发送到服务端
		_, err := c.Write([]byte(content))
		if err != nil {
			log.Println(err)
		}

		n, _ := c.Read(buf)
		log.Println("接收到消息:", string(buf[:n]))
	}
}
