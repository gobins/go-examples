package main

import (
	"fmt"
)

type paper struct {
	length, breadth, thickness int
}

func main() {

	fmt.Printf("String formating: %s\n", "Like this")
	fmt.Printf("Binary printing: %b\n", 10)
	fmt.Printf("Printing Integer: %d\n", 100)
	fmt.Printf("Printing Float: %f\n", 1.23456789)
	fmt.Printf("Printing Hex: %x\n", 100)

	p := paper{1, 2, 3}

	fmt.Printf("Struct printing: %v\n", p)
	fmt.Printf("Struct printing with keys: %+v\n", p)
	fmt.Printf("Struct printing with code point: %#v\n", p)

	fmt.Printf("Printing type of a value: %T\n", p)
}
