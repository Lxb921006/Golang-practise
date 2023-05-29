package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	//m := map[string]interface{}{
	//	"name":  "lxb",
	//	"age":   30,
	//	"skill": "篮球",
	//}

	rd := redis.NewClient(&redis.Options{
		Addr:     "43.156.170.122:6378",
		DB:       10,
		Password: "chatai",
	})

	//rd.HMSet("username5", m)

	data, _ := rd.HGet("prcessstatus", "running").Result()

	fmt.Println("data = ", data)

}
