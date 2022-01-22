package models

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p Product) Print() {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//write a function "ApplyDiscount" that will apply the given discount on the given product
func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * (1 - discount)
}
