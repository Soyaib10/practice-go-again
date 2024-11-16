package main

import "fmt"

type person struct {
	Name string
	Age int
}

// NewPerson can be called factory function/ instance function
func NewPerson(name string, age int) *person {
	return &person{
		Name: name,
		Age: age,
	}
}

func main() {
	newPerson := NewPerson("ami", 12)
	fmt.Println(newPerson.Age)
}