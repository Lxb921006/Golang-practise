package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	//练习2
	wr := make(chan string, 10)
	flag := false
	for i := 1; i <= 10; i++ {
		path := fmt.Sprintf("C:/Users/Administrator/Desktop/111/test%d.txt", i)
		go WriteChan(path, wr)
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go SortData(wr, &flag)
	}
	fmt.Println("11111")
	WhetherCloseChan(wr, &flag)
	wg.Wait()
	fmt.Println("finished")
}

func WriteChan(f string, wr chan string) {
	f1, _ := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	defer f1.Close()
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 1000; i++ {
		f1.WriteString(strconv.Itoa(rd.Intn(1000)) + "\n")
	}
	wr <- f
}

func WhetherCloseChan(wr chan string, fg *bool) {
	for {
		if len(wr) == 10 {
			close(wr)
			*fg = true
			break
		}
	}
}

func SortData(wr chan string, fg *bool) {
	defer wg.Done()
	for {
		if *fg { //这里的意思是先写完,然后再去读
			for {
				v, ok := <-wr
				if !ok {
					break
				}
				ReadData(v)
			}
			break
		}
	}

}

func NewWrite(f string, c []string) {
	f1, _ := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	defer f1.Close()
	for _, v := range c {
		f1.WriteString(v + "\n")
	}
}

func ReadData(f string) {
	c := []string{}
	f1, _ := os.OpenFile(f, os.O_RDONLY, 0777)
	defer f1.Close()
	reader := bufio.NewReader(f1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		newStr := strings.ReplaceAll(str, "\n", "")
		c = append(c, newStr)
	}
	newc := Sort2(c)
	NewWrite(f, newc)

}

func Sort2(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		for t := 0; t < len(s); t++ {
			if t < len(s)-1 {
				numA, _ := strconv.Atoi(s[t])
				numB, _ := strconv.Atoi(s[t+1])
				if numA > numB {
					s[t], s[t+1] = s[t+1], s[t]
				}
			}
		}
	}
	return s
}
