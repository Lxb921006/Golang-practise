package main

import "fmt"

func main() {
	var t interface{}
	i := 3

	t = i

	fmt.Println("t = ", t)

}
