package main

import (
	"fmt"
	"os"
)

type paper struct {
	length, breadth, thickness int
}

func main() {

	fmt.Printf("String formating: %s\n", "Like this")
	fmt.Printf("Binary printing: %b\n", 10)
	fmt.Printf("Character printing: %c\n", 100)
	fmt.Printf("Printing Integer: %d\n", 100)
	fmt.Printf("Scientific notation: %E\n", 123456789.0)
	fmt.Printf("2nd Scientific notation: %e\n", 123456789.0)
	fmt.Printf("Printing Float: %f\n", 1.23456789)
	fmt.Printf("Printing Hex: %x\n", 100)
	fmt.Printf("Quoted Strings: %q\n", "\"Quoted String\"")

	fmt.Printf("|%15s|%15s|\n", "Green", "Lantern")
	fmt.Printf("|%-6d|%-6d|\n", 100, 200)

	p := paper{1, 2, 3}

	fmt.Printf("Struct printing: %v\n", p)
	fmt.Printf("Struct printing with keys: %+v\n", p)
	fmt.Printf("Struct printing with code point: %#v\n", p)

	fmt.Printf("Printing type of a value: %T\n", p)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
