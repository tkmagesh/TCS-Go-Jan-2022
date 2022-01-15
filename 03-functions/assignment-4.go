package main

import "fmt"

func main() {
	increment, decrement := getCounterOperations()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}

func getCounterOperations() (func() int, func() int) { // step - 1
	var counter int           // step - 2
	increment := func() int { // step - 3
		counter++ // step - 4
		return counter
	}

	decrement := func() int { // step - 3
		counter-- // step - 4
		return counter
	}
	return increment, decrement // step - 5
}
