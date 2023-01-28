package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	Project []string `json:"project"`
	Limit   int      `json:"limit"`
}

// 内网循环svn update
func main() {

	rand.Seed(time.Now().Unix())

	var block chan struct{}
	var config Config
	work := make(chan string)

	of, err := os.Open("./projects.json")
	if err != nil {
		log.Print("projects.json not exists, esg = ", err)
		return
	}

	b, err := io.ReadAll(of)
	if err != nil {
		log.Print("read projects.json file, esg = ", err)
		return
	}

	if err = json.Unmarshal(b, &config); err != nil {
		log.Print("failed to parse projects.json, esg = ", err)
		return
	}

	go func() {
		for {
			for _, v := range config.Project {
				work <- v
			}
			time.Sleep(time.Duration(rand.Intn(8)+1) * time.Second)
		}
	}()

	for i := 0; i < config.Limit; i++ {
		go func() {
			for v := range work {
				if err := cmd(v); err != nil {
					log.Printf("%s update failed, esg = %v", v, err)
				}
			}
		}()
	}

	<-block
}

func cmd(p string) (err error) {
	out, err := exec.Command("sh", "/root/shellscript/svn_update2.sh", p).Output()
	if err != nil {
		return errors.New(string(out))
	}
	return
}
