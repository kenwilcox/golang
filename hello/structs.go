package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println("v1 = ", reflect.TypeOf(v1))
	fmt.Println("p = ", reflect.TypeOf(p))
}
