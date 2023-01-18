package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
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
	var block chan struct{}
	var config Config
	work := make(chan string)

	file := "./projects.json"

	of, err := os.Open(file)
	if err != nil {
		log.Print("projects.json not exists, esg = ", err)
		return
	}

	b, err := io.ReadAll(of)
	if err != nil {
		log.Print("read projects.json file, esg = ", err)
		return
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Print("failed to parse projects.json, esg = ", err)
		return
	}

	go func() {
		for {
			for _, v := range config.Project {
				work <- v
			}
			time.Sleep(time.Duration(2) * time.Second)
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
