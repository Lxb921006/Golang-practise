package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	rd := redis.NewClient(&redis.Options{
		Addr: "101.35.143.86:6378",
		DB:   0,
	})

	BrowseGoods(rd)

}

func BrowseGoods(rd *redis.Client) {
	defer rd.Close()
	goods := []string{"篮球", "足球", "羽毛球", "排球", "大球", "浏览记录"}

	id := 0

	for {
		fmt.Println("------------------")
		for i, v := range goods {
			fmt.Println(i, v)
		}
		fmt.Printf("需要看哪个商品输入id就行:")
		fmt.Scanln(&id)
		if id < 0 || id > len(goods)-1 {
			fmt.Println("请输入正确的商品编号")
			continue
		}

		if goods[id] == "浏览记录" {
			data := rd.LRange("goods", 0, 2).Val()
			if len(data) == 0 {
				fmt.Println("你还没有逛任何商品呢")
				continue
			} else {
				fmt.Println("你最近的浏览记录")
				for _, v := range data {
					fmt.Println(v)
				}
			}
		} else {
			rd.LPush("goods", goods[id])
		}
	}
}
