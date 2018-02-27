package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 43, 5, 12, 19}
	c := make(chan int)
	half := len(s) / 2
	go sum(s[:half], c)
	go sum(s[half:], c)
	// by default channel send and receives block until
	// the other side is ready, this allows goroutines
	// to synchronize without explicit locks
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
