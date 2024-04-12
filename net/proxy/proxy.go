package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
)

type ProxyBackend struct {
	Listen  []map[string]string `json:"listen"`
	Backend []map[string]string `json:"backend"`
}

func closeNet(clientConn, remoteConn net.Conn) {
	remoteConn.Close()
	clientConn.Close()
}

func handleClient(clientConn net.Conn, server string, servers []map[string]string) {
	log.Println("recv connect from ", clientConn.RemoteAddr())

	var remoteAddr string

	for _, v1 := range servers {
		sn, ok := v1[server]
		if ok {
			remoteAddr = sn
			break
		}
	}

	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Println("backend connect err >>> ", err)
		return
	}

	go func() {
		defer closeNet(clientConn, remoteConn)
		for {
			// 读取消息长度
			var msgLen uint32
			err := binary.Read(clientConn, binary.LittleEndian, &msgLen)
			if err == io.EOF {
				log.Println(clientConn.RemoteAddr().String(), " client already disconnect")
				return
			}
		}
	}()

	go func() {
		_, err = io.Copy(remoteConn, clientConn)
		if err != nil {
			log.Println("remoteConn to clientConn err >>> ", err)
		}
	}()

	_, err = io.Copy(clientConn, remoteConn)
	if err != nil {
		log.Println("clientConn to remoteConn err >>> ", err)
	}

	log.Printf("proxy to %s finished\n", remoteAddr)
}

func main() {
	var pB ProxyBackend
	file, err := os.ReadFile("servers.json")
	if err != nil {
		log.Fatalln("servers.json not found")
	}

	if err := json.Unmarshal(file, &pB); err != nil {
		log.Fatalln("Unmarshal err >>> ", err)
	}

	if len(pB.Listen) == 0 || len(pB.Backend) == 0 {
		log.Fatalln("json data nil")
	}

	var stop = make(chan struct{})
	for _, servers := range pB.Listen {
		for server, addr := range servers {
			log.Printf("start listen: %s   %s\n", server, addr)
			listener, err := net.Listen("tcp", addr) // 本地监听端口
			if err != nil {
				panic(err)
			}

			defer listener.Close()

			go func(server string) {
				for {
					clientConn, err := listener.Accept()
					if err != nil {
						panic(err)
					}
					go handleClient(clientConn, server, pB.Backend)
				}
			}(server)
		}
	}
	<-stop
}
