package main

import (
	"fmt"
)

func main() {
	a := 5
	var p *int = &a		
	fmt.Print(*p, a)
}
