package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type CountString struct {
	EnCount    int
	NumCount   int
	SpaceCount int
	CharCount  int
	OtherCount int
}

func main() {
	//打开文件
	file := "C:/Users/Administrator/Desktop/test1.txt"
	ff, err := os.OpenFile(file, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ff.Close()

	count := CountString{}
	newff := bufio.NewReader(ff)
	letter := regexp.MustCompile(`[aA-zZ]`)
	num := regexp.MustCompile(`[0-9]`)
	space := regexp.MustCompile(`[\t ]+`)
	china1 := regexp.MustCompile(`\S`)
	china2 := regexp.MustCompile(`\W`)
	for {
		fc, err := newff.ReadString('\n') //读取到\n就换下一行
		if err == io.EOF {
			break
		}
		for _, v := range fc { //可用switch,switch可以不用添加表达式如：switch {...}
			if letter.MatchString(string(v)) {
				count.EnCount++
			} else if num.MatchString(string(v)) {
				count.NumCount++
			} else if space.MatchString(string(v)) {
				count.SpaceCount++
			} else if china1.MatchString(string(v)) && china2.MatchString(string(v)) { //反向匹配中文
				count.CharCount++
			} else {
				count.OtherCount++
			}
		}
	}
	fmt.Println("count.EnCount=", count.EnCount)
	fmt.Println("count.NumCount=", count.NumCount)
	fmt.Println("count.CharCount=", count.CharCount)
	fmt.Println("count.SpaceCount=", count.SpaceCount)
}
