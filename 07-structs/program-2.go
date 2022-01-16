package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//association
type PerishableProduct struct {
	prod   Product
	Expiry string
}

func main() {
	grapes := PerishableProduct{
		prod: Product{
			Id:       501,
			Name:     "Grapes",
			Cost:     10,
			Units:    100,
			Category: "Fruit",
		},
		Expiry: "2 Days",
	}

	fmt.Println(grapes.prod.Name)
	Print(grapes.prod)
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&grapes.prod, 0.1)
	Print(grapes.prod)
}

func Print(p Product) {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//write a function "ApplyDiscount" that will apply the given discount on the given product
func ApplyDiscount(p *Product, discount float64) {
	p.Cost = p.Cost * (1 - discount)
}
