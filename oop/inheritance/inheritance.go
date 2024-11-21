package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Base struct
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "is making a sound!")
}

// Derived struct
type Dog struct {
    Animal
    Breed string
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "says Woof!")
}

func main() {
    var speaker Speaker

    // Base instance
    animal := Animal{Name: "Some Animal"}
    speaker = animal
    speaker.Speak()

    // Derived instance
    dog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    speaker = dog
    speaker.Speak()
}
