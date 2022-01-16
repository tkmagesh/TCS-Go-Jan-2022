package main

import "fmt"

func main() {
	var no int = 10
	/* noPtr  => a variable used to store the "address" of an int value */
	var noPtr *int = &no
	fmt.Println(noPtr)

	//deferencing => accessing the value using the address (pointer)
	var x = *noPtr
	fmt.Println(x)

	//incrementing no
	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	//assignment:6
	var n1, n2 = 10, 20
	fmt.Println("Before swapping n1 = ", n1, " n2 = ", n2)
	swap(&n1, &n2)
	fmt.Println("After swapping n1 = ", n1, " n2 = ", n2)
}

func increment(n *int) {
	(*n)++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
