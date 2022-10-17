package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {
	Client()
}

func Client() {
	//tcp编程客户端
	c, e1 := net.Dial("tcp", "192.168.11.188:8092")
	if e1 != nil {
		fmt.Println(e1)
		return
	}

	defer c.Close()

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
		_, e3 := c.Write([]byte(content))
		if e3 != nil {
			fmt.Println(e3)
		}

		_, e3 = c.Write([]byte("lxb2"))
		if e3 != nil {
			fmt.Println(e3)
		}

		n, _ := c.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}
