package main

import "fmt"

func main() {
	/*
		sayHi := func() {
			fmt.Println("Hi")
		}
	*/
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	/* add := func(x, y int) int {
		return x + y
	} */
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	/*
		func(x, y int) {
			fmt.Println(x * y)
		}(10, 20)
	*/

	var result = func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Println("result = ", result)
}

/* func sayHi() {
	fmt.Println("Hi")
} */
