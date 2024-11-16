package main

import (
	"encapsulation"
	"fmt"
)

func main() {
	std := encapsulation.NewStudent("hello", 25, 101)
	fmt.Println(std.Name)
	fmt.Println(std.GetAge())
	

	fmt.Println(std)
}