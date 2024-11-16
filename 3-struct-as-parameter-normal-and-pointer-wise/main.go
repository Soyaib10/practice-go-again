package main

import "fmt"

// Car struct creating
type Car struct {
	Make  string
	Model string
	Year  int
	Color string
}

// Print car by value
func printCarDetailsByValue(c Car) {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Color: %s\n", c.Make, c.Model, c.Year, c.Color)
}

// Print car by pointer
func printCarDetailsByReference(c *Car) {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Color: %s\n", c.Make, c.Model, c.Year, c.Color)
}

// give a new color
func repaintCar(c *Car, newColor string) {
	c.Color = newColor
}

func main() {

	car := Car{
		Make:  "Tesla",
		Model: "S",
		Year:  2022,
		Color: "Red",
	}

	var make, model, color string
	var year int

	// Get user input for the Car details
	fmt.Print("Enter the make of the car: ")
	fmt.Scanln(&make)

	fmt.Print("Enter the model of the car: ")
	fmt.Scanln(&model)

	fmt.Print("Enter the year of the car: ")
	fmt.Scanln(&year)

	fmt.Print("Enter the color of the car: ")
	fmt.Scanln(&color)

	// making Car instance
	carDynamic := Car{
		Make:  make,
		Model: model,
		Year:  year,
		Color: color,
	}

	printCarDetailsByValue(car) // call by value
	printCarDetailsByReference(&car) // call by address
	repaintCar(&car, "Black") // call by address
	printCarDetailsByReference(&car)

	// Dynamic input using carDynamic- an instance of Car struct
	printCarDetailsByReference(&carDynamic)
	var newColor string
	fmt.Print("Enter a new color for the car: ")
	fmt.Scanln(&newColor)
	repaintCar(&car, newColor)

}
