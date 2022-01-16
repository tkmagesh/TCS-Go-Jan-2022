/* errors */
package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[@main] Recovering from panic")
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Program execution successful")
	}()
	quotient, remainder, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(quotient, remainder)
	fmt.Println("exiting from main")
}

func divide(x, y int) (quotient, remainder int, err error) {
	/* defer func() {
		if r := recover(); r != nil {
			fmt.Println("[@divide] Recovering from panic")
			err = fmt.Errorf("%v", r)
		}
	}() */
	if y == 0 {
		panic("Division by zero")
	}
	quotient = x / y
	remainder = x % y
	return
}
