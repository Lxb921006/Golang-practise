package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdPool *redis.Client

//在主函数(也就是main())执行之前就已经执行
func init() { //初始化
	rdPool = redis.NewClient(&redis.Options{
		Addr:         "101.35.143.86:6377",
		DB:           0,
		MinIdleConns: 5,
		PoolSize:     5,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}

func main() {
	//redis连接池使用

	con := rdPool.Options().PoolSize

	fmt.Println("con=", con)

	s, e := rdPool.Set("name", "lqm", 30*time.Second).Result()

	fmt.Println("s = ", s)

	if e != nil {
		fmt.Println("e = ", e)
	}

	fmt.Println("data = ", rdPool.Get("name").Val())

	defer rdPool.Close()
}
