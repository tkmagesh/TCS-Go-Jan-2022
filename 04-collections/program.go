package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating using index")
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	fmt.Println("Iterating using range")
	for i, v := range nos {
		fmt.Println(i, v)
	}

	fmt.Println("Creating a copy of an array")
	var copyNos = nos
	copyNos[0] = 100
	fmt.Println(nos, copyNos)
	fmt.Printf("Address of nos = %p, copyNos = %p\n", &nos, &copyNos)
}
