package main

import "fmt"

func main() {
	add(100, 200)
	subtract(100, 200)

	logOper(10, 20, add)
	logOper(10, 20, subtract)
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

func logOper(x, y int, oper func(int, int) int) {
	fmt.Println("Before invocation")
	result := oper(x, y)
	fmt.Println("result = ", result)
	fmt.Println("After invocation")
}
