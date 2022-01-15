package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	addWithLog := getLoggedOper(add)
	subtractWithLog := getLoggedOper(subtract)
	addWithLog(100, 200)
	subtractWithLog(100, 200)

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

func getLoggedOper(oper func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("result = ", result)
		fmt.Println("After invocation")
	}
}
