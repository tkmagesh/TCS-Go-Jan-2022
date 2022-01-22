package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/* Stringer interface */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

/* Stringer interface */
func (products Products) String() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

/* sort.Interface interface implementation */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

/* Sorting by name */
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

/* Sorting by cost */
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

/* Sorting by units */
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Intial List")
	fmt.Println(products)

	fmt.Println("Default sort (id)")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("sort by name")
	sort.Sort(ByName{products})
	fmt.Println(products)

	fmt.Println("sort by cost")
	sort.Sort(ByCost{products})
	fmt.Println(products)

	fmt.Println("sort by units")
	sort.Sort(ByUnits{products})
	fmt.Println(products)

}
