package main

import "fmt"

func main() {
	var userChoice, no1, no2, result int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 4 {
			fmt.Println("Invalid choice")
			continue
		}
		no1, no2 = getOperands()
		result = process(userChoice, no1, no2)
		fmt.Printf("result = %d\n", result)
	}
}

func process(userChoice, no1, no2 int) int {
	var result int
	switch userChoice {
	case 1:
		result = add(no1, no2)
	case 2:
		result = subtract(no1, no2)
	case 3:
		result = multiply(no1, no2)
	case 4:
		result = divide(no1, no2)
	}
	return result
}
func getOperands() (int, int) {
	var no1, no2 int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d %d", &no1, &no2)
	return no1, no2
}
func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice :")
	fmt.Scanf("%d", &userChoice)
	return userChoice
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
