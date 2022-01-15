package main

import "fmt"

func main() {
	var userChoice, no1, no2, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice :")
		fmt.Scanf("%d", &userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 4 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the number :")
		fmt.Scanf("%d %d", &no1, &no2)
		switch userChoice {
		case 1:
			result = no1 + no2
		case 2:
			result = no1 - no2
		case 3:
			result = no1 * no2
		case 4:
			result = no1 / no2
		}
		fmt.Printf("result = %d\n", result)

	}
}
