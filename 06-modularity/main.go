package main

import (
	"fmt"
	"modularity-demo/calc"
	"modularity-demo/calc/utils"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Operation Count : ", calc.GetOpCount())

	fmt.Println("Is 97 prime? :", utils.IsPrime(97))
}
