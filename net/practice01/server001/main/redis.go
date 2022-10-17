package main

import (
	"time"

	"github.com/go-redis/redis"
)

var RdPool *redis.Client

//初始化redis连接池
func InitPoolRds(addr string) {
	RdPool = redis.NewClient(&redis.Options{
		Addr:         addr,
		DB:           0,
		MinIdleConns: 5,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
}
