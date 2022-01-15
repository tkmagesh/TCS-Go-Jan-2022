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

	//switch construct
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Invalid Score"
	*/

	score := 6
	/*
		switch score {
		case 0:
		case 1:
		case 2:
		case 3:
			fmt.Println("Terrible")
		case 4:
		case 5:
		case 6:
			fmt.Println("Not Bad")
		case 7:
		case 8:
		case 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6:
		fmt.Println("Not Bad")
	case 7, 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	myVar := 22
	/*
		switch myVar % 2 {
		case 0:
			fmt.Println("even")
		case 1:
			fmt.Println("odd")
		}
	*/
	switch {
	case myVar%2 == 0:
		fmt.Println("even")
	case myVar%2 == 1:
		fmt.Println("odd")
	}

	//fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}

}
