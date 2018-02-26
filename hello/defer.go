package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	deferEval(10)
	deferStack()
}

func deferEval(i int) {
	// this is printed at the end, but it uses the value given
	defer fmt.Printf("i = %v\n", i)
	i += 10
	fmt.Printf("defer value = %v\n", i)
}

func deferStack() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		//noinspection GoDeferInLoop
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
