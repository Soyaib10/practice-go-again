package main

import "fmt"

func typeAssertion() {
	var i interface{} = "Hello world!"

	 s := i.(string)
	 fmt.Println(s)

	 n, ok := i.(int)
	 if ok {
		fmt.Println(n)
	 } else {
		fmt.Println("Type assertion failed!")
	 }
}