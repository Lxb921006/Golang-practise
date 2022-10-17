package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	m := map[string]interface{}{
		"name":  "lxb",
		"age":   30,
		"skill": "篮球",
	}

	rd := redis.NewClient(&redis.Options{
		Addr: "101.35.143.86:6378",
		DB:   0,
	})

	rd.HMSet("username5", m)

	data := rd.HMGet("username5", "name", "age", "skill").Val()

	fmt.Println("data = ", data)

}
