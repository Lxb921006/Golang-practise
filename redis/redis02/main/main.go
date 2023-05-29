package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var rdPool *redis.Client

// 在主函数(也就是main())执行之前就已经执行
func init() { //初始化
	rdPool = redis.NewClient(&redis.Options{
		Addr:         "43.156.170.122:6378",
		DB:           10,
		MinIdleConns: 5,
		Password:     "chatai",
		PoolSize:     5,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}

func main() {
	//redis连接池使用

	var data1 = make(map[string]string)

	log.Printf("type = %s", data1["result"])

	//con := rdPool.Options().PoolSize
	//
	//fmt.Println("con=", con)
	//
	//s, e := rdPool.HSet("aaaaaa", "name", "lxb").Result()
	//
	//fmt.Println("s = ", s)

	result, err := rdPool.HGet("prcessstatus", "running").Result()
	if err != nil {
		log.Println(err)
		return
	}

	data1["result"] = result

	log.Println(data1)

	defer rdPool.Close()
}
