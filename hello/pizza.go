package main

import (
	"time"
	"strconv"
	"fmt"
)

var pizzaNumber = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNumber ++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNumber)
	fmt.Println("Make Dough and Send for Sauce")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <- stringChan
	fmt.Println("Add Sauce and Send", pizza, "for toppings")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	pizza := <- stringChan
	fmt.Println("Add Toppings to", pizza, "and ship")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func main() {
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 1000)
	}
}
