package main

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	close(c)
	close(c)
	// c <- 3

}
