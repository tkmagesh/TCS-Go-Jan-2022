package main

import "fmt"

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

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(p Product) bool {
	return products.IndexOf(p) >= 0
}

/*
func (products Products) FilterCostlyProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
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

	marker := Product{103, "Marker", 50, 20, "Utencil"}
	dummy := Product{5000, "Dummy", 50, 20, "Dummy"}

	//IndexOf
	fmt.Println("Index of Marker = ", products.IndexOf(marker))
	fmt.Println("Index of Dummy = ", products.IndexOf(dummy))

	//Includes
	fmt.Println("Is Marker included ? = ", products.Includes(marker))
	fmt.Println("Is Dummy  included ? = ", products.Includes(dummy))

	/*
		//Filter Costly Products
		fmt.Println()
		costlyProducts := products.FilterCostlyProducts()
		fmt.Println("Costly Products")
		fmt.Println(costlyProducts.ToString())

		//Filter Stationary Products
		fmt.Println()
		stationaryProducts := products.FilterStationaryProducts()
		fmt.Println("Stationary Products")
		fmt.Println(stationaryProducts.ToString())
	*/

	//Filter Costly Products
	fmt.Println()
	costlyProductCriteria := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductCriteria)
	fmt.Println("Costly Products")
	fmt.Println(costlyProducts)

	//Filter Stationary Products
	fmt.Println("Filter")
	fmt.Println()
	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductCriteria)
	fmt.Println("Stationary Products")
	fmt.Println(stationaryProducts)

	fmt.Println("Any")
	fmt.Println("Are there any costly products ? : ", products.Any(costlyProductCriteria))
	fmt.Println("Are there any stationary products ? : ", products.Any(stationaryProductCriteria))

	fmt.Println("All")
	fmt.Println("Are all costly products ? : ", products.All(costlyProductCriteria))
	fmt.Println("Are all stationary products ? : ", products.All(stationaryProductCriteria))

}

/*
Write the following methods
IndexOf => return the index of the given product in the collection
Includes => return true if the given product is present in the collection else return false
Filter => return products that matches the given criteria
	Use cases:
		1. filter all costly products (cost > 1000)
		2. filter all stationary products (category == stationary)
All => return true if ALL the products matches the given criteria else return false
	Use cases:
		1. Are all the products costly products (cost > 1000) ?
		2. Are all the products stationary products? (category == "Stationary")

Any => return true if ANY of the products matches the given criteria else returns false
	Use cases:
		1. Are there ANY costly product (cost > 1000) ?
		2. Are there ANY stationary products? (category == "Stationary")

*/
