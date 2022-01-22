package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main started")
	go f1()
	f2()

	/* DO NOT DO THIS */
	//time.Sleep(500 * time.Millisecond)

	/* DO NOT DO THIS */

	fmt.Println("hit ENTER to continue")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Main completed")
}

func f1() {
	//time.Sleep(2 * time.Second)
	fmt.Println("f1 is invoked")
}

func f2() {
	fmt.Println("f2 is invoked")
}
