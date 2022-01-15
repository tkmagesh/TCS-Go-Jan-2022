package main

import "fmt"

func main() {
	/*
		var x int
		x = 42
		fmt.Println(x)
	*/

	/*
		var x int = 42
		fmt.Println(x)
	*/

	/*
		var x = 42
		//fmt.Println(x)
		fmt.Printf("x (value=%d) is of type %T\n", x, x)
	*/

	/*
		x := 42 // => this syntax is allowed ONLY in a function
		fmt.Printf("x (value=%d) is of type %T\n", x, x)
	*/

	//Hanling multiple variables
	/*
		var x int
		var y int
		var result int
		x = 10
		y = 20
		result = x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	/*
		var x, y, result int
		x = 10
		y = 20
		result = x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	/*
		var x, y, result int
		x, y = 10, 20
		result = x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	/*
		var x, y int = 10, 20
		var result int = x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	/*
		var x, y = 10, 20
		var result = x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	/*
		x, y := 10, 20
		result := x + y
		fmt.Printf("%d + %d = %d\n", x, y, result)
	*/

	var (
		x, y   = 10, 20
		result = x + y
	)
	fmt.Printf("%d + %d = %d\n", x, y, result)
}
