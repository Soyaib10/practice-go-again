package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

// Nested Stucts
type Address struct {
	Street string
	City   string
	State  string
}

type PersonPerson struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// Struct with methods
func (p PersonPerson) Greet() string {
	return "Hello, my name is " + p.Name + " from PersonPerson struct."
}

// mehtod overriding. using Greet() as method, and Address & PersonPerson as struct.

func (a Address) Greet() string {
	return "Hi, I am " + a.City + " from Address struct."
}

// Constructors: While Go doesn’t have constructors in the same way as some other languages, you can define factory functions:

func constructorsInGo(name string, age int, email string) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func main() {

	// populate the struct
	person1 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@gmail.com",
	}
	fmt.Println(person1.Name)
	fmt.Printf("%+v\n", person1)

	// populate struct using pointer
	person2 := &Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@gmail.com",
	}
	fmt.Println(person2.Name)

	person := PersonPerson{
		Name:  "Alice",
		Age:   24,
		Email: "alice@example.com",
		Address: Address{
			Street: "1/2 lincon street",
			City:   "Jamalpur",
			State:  "CA",
		},
	}

	newPerson := PersonPerson{
		Name:  "Alice",
		Age:   24,
		Email: "alice@example.com",
		Address: Address{
			Street: "1/2 lincon street",
			City:   "Jamalpur",
			State:  "CA",
		},
	}

	newAddress := Address{
		Street: "ami jani na",
		City:   "taw jani na",
		State:  "bollam na jani na",
	}

	fmt.Println(person.Name)
	fmt.Println(person.Address.Street)

	// call methods
	fmt.Println(person.Greet())
	fmt.Println(newPerson.Greet())

	// method overloading
	fmt.Println(newAddress.Greet()) // from Address struct
	fmt.Println(newPerson.Greet())  // from PersonPerson struct

	// object for constructorsInGo
	constructorsInGoObject := constructorsInGo("alice", 22, "af@mail.com")
	fmt.Println(constructorsInGoObject.Age)
}
