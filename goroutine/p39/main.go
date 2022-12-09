package main

import (
	"log"
	"math/rand"
	"time"
)

//池化技术: 即使后面有大量客户等待都不会再创建goroutine，因为已经提前创建了10个goroutine

func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ customer#", c, " drinks at the bar")
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- customer#", c, " leaves the bar")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const BarSeatCount = 10
	consumers := make(chan int)

	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		// time.Sleep(time.Second)
		consumers <- customerId
	}
}
