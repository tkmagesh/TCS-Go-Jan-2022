package main

import "fmt"

func main() {
	process(10, 20, add)
	process(10, 20, subtract)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func process(x, y int, oper func(int, int) int) {
	result := oper(x, y)
	fmt.Println(result)
}
