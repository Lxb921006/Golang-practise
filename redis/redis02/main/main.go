package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdPool *redis.Client

// 在主函数(也就是main())执行之前就已经执行
func init() { //初始化
	rdPool = redis.NewClient(&redis.Options{
		Addr:         "43.153.55.148:6377",
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
	time.Now().Format(time.DateTime)
	setQuota()
}

func setInvite() {
	var initData = []string{"lxb"}
	var data = map[string][]string{
		"user": initData,
	}
	b, err := json.Marshal(&data)
	err = rdPool.HSet("invite", "ogR3E62jXXJMbVcImRqMA1gTSegM", b).Err()
	fmt.Println("setInvite HSet err >>>", err)

	str, err := rdPool.HGet("invite", "ogR3E62jXXJMbVcImRqMA1gTSegM").Result()
	fmt.Println("setInvite HGet err >>>", str, err)

	defer rdPool.Close()
}

func setQuota() {
	ds, err := rdPool.Get("openia").Result()
	fmt.Println("ds >>> ", ds, ", err >>>", err)

	var data = map[string]int{
		"chatgpt":  976,
		"qw":       976,
		"gemini":   960,
		"bd":       972,
		"time":     1706174482,
		"add":      2,
		"invite":   0,
		"finished": 2,
	}

	b, _ := json.Marshal(&data)

	err = rdPool.HSet("quota", "ogR3E62jXXJMbVcImRqMA1gTSegM", b).Err()
	fmt.Println("err >>>", err)

	var di = make(map[string]int, 1)
	b1, err := rdPool.HGet("quota", "ogR3E62jXXJMbVcImRqMA1gTSegM").Result()
	_ = json.Unmarshal([]byte(b1), &di)
	fmt.Println("di >>> ", di, ", err >>>", err)

	k, err := rdPool.HExists("quota", "ogR3E62jXXJMbVcImRqMA1gTSegM").Result()
	fmt.Println("k >>> ", k, ", err >>>", err)

	defer rdPool.Close()
}
