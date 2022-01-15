package main

import "fmt"

func main() {
	fmt.Println(sum("abc", 1, 2, 3, 4, 5))
}

func sum(s string, nos ...int) int {
	//nos => array of ints
	var total int
	for i := 0; i < len(nos); i++ {
		total += nos[i]
	}
	return total
}
