package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		rd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3)
		fmt.Println(rd)
		time.Sleep(time.Duration(rd) * time.Second) //随机休眠
	}
}
