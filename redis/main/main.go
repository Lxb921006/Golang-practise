package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type UserInfo struct {
	UserId  int
	UserPwd string
	Content string
}

// type userList []UserInfo

func main() {
	//redis使用
	rd := redis.NewClient(&redis.Options{
		Addr:     "43.156.18.45:6378",
		DB:       4,
		Password: "chatai",
	})

	result, err2 := rd.Ping().Result()
	fmt.Printf("result: %v, err2: %v", result, err2)

	defer rd.Close()

	//rd.Del("username").Val()
	//
	//rd.Del("username4").Val()
	//
	//res := rd.LPush("username", "lxb", "lqm", "lyy", "lch").Val()
	//
	//fmt.Println("res=", res)
	//
	//fmt.Println(rd.LRange("username", 0, -1).Val())
	//
	//// rd.HSet("username2", "name", "lxb")
	//
	//// m := map[string]interface{}{
	//// 	"age":   30,
	//// 	"hobby": []string{"篮球", "足球"},
	//// }
	//
	//res2 := rd.HMSet("username4", map[string]interface{}{
	//	"age":   30,
	//	"hobby": "篮球",
	//})
	//
	//fmt.Println("res2 = ", res2)
	//
	//d, err := rd.HMGet("username4", "age").Result()
	//fmt.Println("d = ", d[0].(string))
	//fmt.Println("errd = ", err)
	//
	//// b, err := rd.Exists("username111").Result()
	//
	//b, err := rd.Keys("username4").Result()
	//
	//fmt.Println("b = ", b)
	//fmt.Println("berr = ", err)
	//
	//res3, _ := rd.HGetAll("username4").Result()
	//fmt.Println("rd.HGetAll = ", res3)

	// fmt.Println("username4=", rd.HMGet("username4", "age").Val())

	// res3, _ := rd.HGetAll("users").Result()
	// fmt.Println("rd.HGetAll = ", res3)

	// for i, v := range res3 {
	// 	fmt.Printf("uid=%v, userinfo=%v\n", i, v)
	// }

	// res1, err1 := rd.HGet("users", "ages").Result()

	// if err1 != nil {
	// 	fmt.Println("err1 = ", err1)
	// 	// return
	// }
	// fmt.Println("res1 = ", res1)

	// var u userList

	// for i := 0; i < 3; i++ {
	// 	ui := UserInfo{
	// 		UserId:  100,
	// 		UserPwd: "",
	// 		Content: "aaa",
	// 	}
	// 	u = append(u, ui)
	// }
	// fmt.Println("u = ", u)

	// dd := map[int]interface{}{
	// 	100: u,
	// }

	// data, _ := json.Marshal(&dd)

	// _, e := rd.HSet("chat", "100", string(data)).Result()
	// if e != nil {
	// 	fmt.Println("写入redis失败")
	// }

	// s, e := rd.HGet("chat", "999").Result()
	// if e != nil {
	// 	fmt.Println("e = ", e)
	// } else {
	// 	fmt.Println("s = ", s)
	// }

}
