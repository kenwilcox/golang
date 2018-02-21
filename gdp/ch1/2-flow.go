package main

import "fmt"

func main() {
	// if...else
	ten := 10
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else {
		println("Ten is not equal to 20")
	}

	// else...if
	if "a" == "b" || 10 == 10 || true == false {
		println("10 is equal to 10")
	} else if 11 == 11 && "go" == "go" {
		println("This isn't printed because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}

	// switch
	number := 3
	switch number {
	case 1:
		println("Number is 1")
	case 2:
		println("Number is 2")
	case 3:
		println("Number is 3")
	}

	// for
	for i := 0; i <= 10; i++ {
		println(i)
	}

	// for...range
	var myArray = [...]int{1, 2, 3, 4, 5, 6}
	for index, value := range myArray {
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}

	// range is just this
	for index := 0; index < len(myArray); index++ {
		value := myArray[index]
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}
}
