package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	//var p Product
	//p := Product{}
	//p := Product{101, "Pen", 10, 100, "Stationary"}
	p := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}

	//fmt.Println(p)
	Print(p)

	//apply 10% discount on p and print it
}

func Print(p Product) {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//write a function "ApplyDiscount" that will apply the given discount on the given product
