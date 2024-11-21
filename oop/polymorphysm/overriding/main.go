package main

import "fmt"

type Car struct {
	Name  string
	Model string
}

type Truck struct {
	Name  string
	Model string
}

func (c Car) Sound() string {
	return "sound of Car"
}

func (t Truck) Sound() string {
	return "sound of Truck"
}

func main() {
	car := Car{Name: "c", Model: "cc"}
	truck := Truck{Name: "t", Model: "tt"}

	fmt.Println(car)
	fmt.Println(car.Sound())
	fmt.Println(truck)
	fmt.Println(truck.Sound())
}
