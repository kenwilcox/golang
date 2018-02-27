package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// even an unused receiver makes it not panic
	f, _ = i.(float64)
	fmt.Println(f)

	f = i.(float64) // panic
	fmt.Println(f)
}
