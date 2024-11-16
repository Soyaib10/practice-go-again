package main

import "fmt"

func array() {
	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = i + 1
	}
	slice = append(slice, 4, 5)
	fmt.Println(slice)
}
