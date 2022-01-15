package main

import "fmt"

func main() {
	fmt.Println("Enter the number : ")
	var input int
	fmt.Scanf("%d", &input)
	for i := 2; i < input; i++ {
		if input%i == 0 {
			fmt.Println("Not Prime")
			return
		}
	}
	fmt.Println("Prime")
}
