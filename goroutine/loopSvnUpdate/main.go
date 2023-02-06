package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// svn update
type Config struct {
	Project []string `json:"project"`
	Limit   int      `json:"limit"`
	TimeOut int      `json:"timeout"`
}

func (c *Config) PareJson(file string) (b []byte, err error) {
	of, err := os.Open(file)
	if err != nil {
		return
	}

	b, err = io.ReadAll(of)
	if err != nil {
		return
	}

	return
}

func (c *Config) GetFileData(file string) (cc Config, err error) {
	b, err := c.PareJson(file)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &cc); err != nil {
		return
	}

	return
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var block chan struct{}
	var config Config
	work := make(chan string)

	ctx := context.Background()

	config, err := config.GetFileData("./projects.json")
	if err != nil {
		log.Print(err)
		return
	}

	go func() {
		for {
			for _, v := range config.Project {
				work <- v
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
		}
	}()

	for i := 0; i < config.Limit; i++ {
		go func(ctx context.Context) {
			for v := range work {
				if err := cmd(v, ctx, config); err != nil {
					log.Printf("%s update failed, esg = %v", v, err)
				}
			}
		}(ctx)
	}

	<-block
}

func cmd(p string, ctx context.Context, config Config) (err error) {
	if config.TimeOut > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(config.TimeOut))
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, "sh", "/root/shellscript/svn_update2.sh", p)
	if err = cmd.Run(); err != nil {
		return
	}

	return

}
