package main

import "fmt"

func main() {
	sayHi := getSayHi()
	sayHi()

	add := getAdder()
	fmt.Println(add(100, 200))
}

func getSayHi() func() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	return sayHi
}

func getAdder() func(int, int) int {
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	return add
}
