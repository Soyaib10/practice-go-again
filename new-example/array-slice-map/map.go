package main

import "fmt"

func maps() {
	// declaration
	var myMap map[string]int
	mMap := make(map[string]int)

	myMap = map[string]int{
		"one": 1,
		"two": 2,
	}

	mMap = map[string]int{
		"one": 1,
		"two": 2,
	}

	mMap["three"] = 3
	value := mMap["three"]
	delete(mMap, "three")

	if _, ok := mMap["two"]; ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println(myMap)
	fmt.Println(mMap)
	fmt.Println(value)
	fmt.Println(mMap)
}