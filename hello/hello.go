package main

import (
	"fmt"
	"math"
	"github.com/kenwilcox/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.Reverse("!nuf si sihT"))

	favNums := [5]float64 {1, 2, 3, 4, 5}

	for i, value := range favNums {
		fmt.Println(value, i)
	}

	// get the first two elements
	fmt.Println("favNums[:2] =", favNums[:2])

	// skip the first two elements
	fmt.Println("favNums[2:] =", favNums[2:])

	fmt.Println(factorial(3))

	defer printTwo()
	printOne()

	fmt.Println("1 / 0 =", safeDiv(1, 0))
	fmt.Println("3 / 2 =", safeDiv(3, 2))

	x := 0
	changeXVal(x)
	fmt.Println("x = ", x)
	changeXValNow(&x)
	fmt.Println("x = ", x)

	rect1 := Rect {leftX: 0, topY: 50, height: 10, width:10}
	rect2 := Rect {0, 50, 10, 10}

	fmt.Println("rect1 and rect1 leftX match: ", rect1.leftX == rect2.leftX)
	fmt.Println("rect1 and rect1 topY match: ", rect1.topY == rect2.topY)
	fmt.Println("rect1 and rect1 height match: ", rect1.height == rect2.height)
	fmt.Println("rect1 and rect1 width match: ", rect1.width == rect2.width)
	fmt.Println("rect1 and rect1 area match: ", rect1.area() == rect2.area())

	// interface
	rect := Rectangle{20, 50}
	circle := Circle{4}
	triangle := Triangle{5, 5, 5,}

	fmt.Println("Rectangle Area = ", getArea(rect))
	fmt.Println("Circle Area = ", getArea(circle))
	fmt.Println("Triangle Area = ", getArea(triangle))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num - 1)
}

func printOne() {fmt.Println("Print One")}
func printTwo() {fmt.Println("Print Two")}

func safeDiv(num1, num2 int) int {

	// This will show the error
	//defer func() {fmt.Println(recover())}()
	// This will not show the error
	defer func() {recover()}()

	return num1 / num2
}

func changeXVal(x int) {
	x = 2
}

func changeXValNow(x *int){
	*x = 2
}

type Rect struct {
	leftX float64
	topY float64
	height float64
	width float64
}

// add a method to Rectangle
func (rect *Rect) area() float64 {
	return rect.width * rect.height
}

// using interfaces
type Shape interface {
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t Triangle) area() float64 {
	s := (t.a + t.b + t.b) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func getArea(shape Shape) float64 {
	return shape.area()
}