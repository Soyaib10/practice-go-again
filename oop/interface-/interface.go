package main

import (
	"fmt"
	"math"
)

// interfaces
type Shape interface {
	Area() float64
}

type Measureable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measureable
}

// Rectange struct
type Rectange struct {
	width, hight float64
}

func (r Rectange) Area() float64 {
	return r.hight * r.width
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.width + r.hight)
}

// Circle Struct
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// implementing 
func calculateArea(s Shape) float64 {
	return s.Area()
}

func describeShape(g Geometry) {
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

// Main function
func main() {
	// calculateArea can print recangle and circle casue they both are Shape interface type
	rectange := Rectange{width: 34.3, hight: 43.3}
	circle := Circle{radius: 4.4}
	fmt.Println(calculateArea(rectange))
	fmt.Println(calculateArea(circle))


	// creating a slice of Shpape interface
	slice := []Shape{rectange, circle}
	fmt.Println(slice[1])
	for _, x := range slice {
		fmt.Println(x.Area())
	}

	// geometry interface
	fmt.Println("\nGeometry interface")
	describeShape(rectange)
}