/* Modify the program so that the add function can be executed as a goroutine */
package main

import "fmt"

func main() {
	result := go add(100, 200)
	fmt.Println("Result = ", result)
}

func add(x, y int) int {
	fmt.Printf("processing %d and %d\n", x, y)
	result := x + y
	fmt.Printf("returing result : ", result)
	return result
}
