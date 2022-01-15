/* errors */
package main

import (
	"errors"
	"fmt"
)

func main() {
	quotient, remainder, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotient, remainder)
	}
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("Cannot divide by zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
