package main

import "fmt"

func main() {
	var a interface{}
	a = "100"
	if val, ok := a.(int); ok {
		fmt.Printf("a is of type int with val = %d\n", val)
	} else {
		fmt.Println("a is not an int")
	}

	//a = true
	//a = 100
	//a = "This is a string"
	//a = []int{10, 20, 30, 40}
	//a = 10.3987
	a = struct{}{}
	switch val := a.(type) {
	case int:
		fmt.Println("a is an int, a + 100 = ", val+100)
	case string:
		fmt.Println("a is string, len(a) = ", len(val))
	case bool:
		fmt.Println("a is bool, a = ", val)
	case []int:
		fmt.Println("a is []int with count = ", len(val))
	default:
		fmt.Println("Unknown type")
	}

}
