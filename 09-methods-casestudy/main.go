package main

import (
	"fmt"
	"methods-demo/models"
)

type MyProduct models.Product

func (p MyProduct) WhoAmI() {
	fmt.Println("I am ", models.Product(p).Name)
}

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	p := MyProduct(models.Product{100, "Pen", 10, 100, "Stationary"}) //typecast Product to MyProduct
	p.WhoAmI()

	s := MyStr("This is a sample string") //type cast the string to MyStr
	fmt.Println(s.Length())
}
