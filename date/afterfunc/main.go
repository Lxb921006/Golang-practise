package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// Defining duration parameter of
	// AfterFunc() method
	DurationOfTime := time.Duration(3) * time.Second

	// Defining function parameter of
	// AfterFunc() method
	f := func() {

		// Printed when its called by the
		// AfterFunc() method in the time
		// stated above
		fmt.Println("Function called by " +
			"AfterFunc() after 3 seconds")
	}

	// Calling AfterFunc() method with its
	// parameter
	time.AfterFunc(DurationOfTime, f)

	// Calling stop method
	// w.r.to Timer1
	// defer Timer1.Stop()

	run()

}

func run() int {

	time.Sleep(time.Second * 6)

	s := []string{"lxb", "lqm"}
	s1 := strings.Join(s, "")

	fmt.Println(10, s1)
	return 10
}
