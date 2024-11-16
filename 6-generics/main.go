package main

import (
	"fmt"
)

// example 1
func add[T interface{}](a, b T) T { // interface{}
	return a
}

// example 2
func add2[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	resultOfAdd := add(1, 3)
	fmt.Println(resultOfAdd)

	resultOfAdd2 := add2(1, 3)
	fmt.Println(resultOfAdd2)
}