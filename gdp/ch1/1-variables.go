package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicit string = "Hello, I'm an explicitly declared variable"
	inferred := "Hello, I'm an inferred variable"

	fmt.Println(explicit)
	fmt.Println(inferred)
	fmt.Println("Varialble 'explicit' is of type:", reflect.TypeOf(explicit))
	fmt.Println("Variable 'inferred' is of type:", reflect.TypeOf(inferred))
}
