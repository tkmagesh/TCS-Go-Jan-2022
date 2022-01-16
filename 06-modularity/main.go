package main

import (
	"fmt"
	"modularity-demo/calc"
	"modularity-demo/calc/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	color.Red(fmt.Sprintln("Operation Count : ", calc.GetOpCount()))

	color.Blue(fmt.Sprintln("Is 97 prime? :", utils.IsPrime(97)))
}
