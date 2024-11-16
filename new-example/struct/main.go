package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

func main() {
	// constructor function
	person := NewPerson("constructor function", 25)
	person.Name = "constructor"
	fmt.Println(person)

	// struct Literal- think like calling constructor using object
	person1 := Person{Name: "struct literal", Age: 25}
	person2 := &Person{Name: "Pointer literal", Age:  45}

	// avoid the rest
	person3 := Person{"positional syntex", 40} 
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	// zero value initialization
	var person4 Person
	person4.Name = "zero value initialization"
	person4.Age = 23
	fmt.Println(person4)

	// new keyword
	person5 := new(Person)
	person5.Name = "new keyword"
	person5.Age = 23
	fmt.Println(person5)
}