package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//练习
	f1 := "C:/Users/Administrator/Desktop/test.txt"
	f2 := "C:/Users/Administrator/Desktop/test2.txt"
	s := "gAxxx"
	s1 := strings.ReplaceAll(s, "g", "")
	fmt.Println(s1)
	fmt.Println(s)
	s2 := "xxx"
	match := strings.Contains(s2, "xxx")
	if match {
		fmt.Println("ok")
	}

	f11, err11 := ioutil.ReadFile(f1)
	if err11 != nil {
		log.Fatal(err11)
		return
	}

	err22 := ioutil.WriteFile(f2, f11, 0777)
	if err22 != nil {
		log.Fatal(err22)
		return
	}
	fmt.Println("写入完成")

}
