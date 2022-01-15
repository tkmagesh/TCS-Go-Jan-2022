package main

import "fmt"

func main() {
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is even")
	} else {
		fmt.Println(no, "is odd")
	}

	//for construct
	//ver 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)

	//ver 2.0
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	//ver 3.0
	//infinite loop
	no := 1
	for {
		if no > 100 {
			break
		}
		no += no
	}
	fmt.Printf("no = %d\n", no)
}
