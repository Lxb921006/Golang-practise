package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

type Customer struct{ id int }
type Bar chan Customer

func (bar Bar) ServeCustomer(c Customer) {
	log.Print("++ customer#", c.id, " starts drinking")
	log.Print("gn = ", runtime.NumGoroutine())
	time.Sleep(time.Second * time.Duration(4+rand.Intn(16)))
	log.Print("-- customer#", c.id, " leaves the bar")
	<-bar // leaves the bar and save a space
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// The bar can serve most 10 customers
	// at the same time.
	bar24x7 := make(Bar, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Millisecond)
		customer := Customer{customerId}
		select {
		case bar24x7 <- customer: // try to enter the bar
			go bar24x7.ServeCustomer(customer)
		default:
			log.Print("customer#", customerId, " goes elsewhere")
		}
	}
}
