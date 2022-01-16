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

	fmt.Println("Slices")
	var products []string
	fmt.Println(products, len(products), cap(products))
	//var products = []string{"Laptop", "Mobile", "Tablet", "Camera"}
	/*
		products = append(products, "Laptop")
		fmt.Println(products, len(products), cap(products))
		products = append(products, "Mobile")
		fmt.Println(products, len(products), cap(products))
		products = append(products, "Tablet")
		fmt.Println(products, len(products), cap(products))
	*/

	products = append(products, "Laptop", "Mobile", "Tablet", "Camera")
	fmt.Println(products, len(products), cap(products))
	products = append(products, "Pen")
	fmt.Println(products, len(products), cap(products))

	var n = make([]int, 5, 10)
	fmt.Println(n, len(n), cap(n))
	/* n = append(n, 11)
	fmt.Println(n, len(n), cap(n)) */

	products = append(products, "Pencil", "Eraser", "Marker", "Scribble-Pad")
	fmt.Println(products)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from start to hi-1
		[:] => copy of the slice
	*/
	fmt.Println("Slicing")
	fmt.Println("products[2:5] = ", products[2:5])
	fmt.Println("products[2:] = ", products[2:])
	fmt.Println("products[:5] = ", products[:5])

	dupProducts := products[:]
	dupProducts[0] = "Dummy"
	fmt.Println(products, dupProducts)
}
