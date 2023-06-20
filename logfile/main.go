package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\my.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	log.Println(" Orders API Called")

	fi, err := os.Stat("C:\\Users\\Administrator\\Desktop\\my.log")
	if err != nil {
		return
	}

	log.Println("size >>>,", fi.Size(), fi.Name())
}
