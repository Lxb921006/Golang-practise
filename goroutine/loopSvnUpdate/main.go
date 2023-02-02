package main

import (
	"context"
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

var (
	stop = make(chan int, 1)
)

// svn update
func main() {

	rand.Seed(time.Now().Unix())
	var block chan struct{}
	work := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	config, err := getFileData("./projects.json")
	if err != nil {
		log.Print(err)
		return
	}

	go func() {
		for {
			for _, v := range config.Project {
				work <- v
			}

			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)

		}
	}()

	for i := 0; i < config.Limit; i++ {
		go func(ctx context.Context) {
			for v := range work {
				if err := cmd(v, ctx); err != nil {
					log.Printf("%s update failed, esg = %v", v, err)
				}
			}
		}(ctx)

		go func(ctx context.Context) {
			select {
			case <-ctx.Done():

				log.Printf("Context cancelled: %v\n", ctx.Err())
			default:
			}
		}(ctx)

	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				stop <- 1
			case <-time.After(2 * time.Second):
				stop <- 1
			}
		}
	}()

	<-block
}

func cmd(p string, ctx context.Context) (err error) {
	for {
		select {
		case <-stop:
			return errors.New("run cmd timeout")
		default:
			//do something
			out, err := exec.Command("sh", "/root/shellscript/test.sh", p).Output()
			if err != nil {
				return errors.New(string(out))
			}
		}
	}
}

func public(file string) (b []byte, err error) {
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

func getFileData(file string) (c Config, err error) {
	b, err := public(file)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &c); err != nil {
		return
	}

	return
}
