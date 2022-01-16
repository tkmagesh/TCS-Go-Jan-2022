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
	p := &Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}

	//fmt.Println(p)
	//Print(p)
	p.Print()

	//apply 10% discount on p and print it
	//ApplyDiscount(&p, 0.1)
	p.ApplyDiscount(0.1)
	//Print(p)
	p.Print()

}

func (p Product) Print() {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//write a function "ApplyDiscount" that will apply the given discount on the given product
func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * (1 - discount)
}
