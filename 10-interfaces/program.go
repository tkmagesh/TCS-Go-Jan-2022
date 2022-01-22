package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

//implicity implementation of ShapeWithArea interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//implicity implementation of ShapeWithPerimter interface
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

//implicity implementation of ShapeWithArea interface
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//implicity implementation of ShapeWithPerimter interface
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

//Utility
type ShapeWithArea interface {
	Area() float64
}

func PrintArea(shapeWithArea ShapeWithArea) {
	fmt.Println("Area = ", shapeWithArea.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(shapeWithPerimeter ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shapeWithPerimeter.Perimeter())
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Println("Area = ", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 12, Width: 10}
	//fmt.Println("Area = ", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}
