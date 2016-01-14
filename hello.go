package main

import (
    "fmt"
    "runtime"
    "reflect"
)

var (
	name = "Gobin"
	course = "Golang"
	module = "Go Fundamentals"
)

func main() {

	// fmt.Println("Hello from", runtime.GOOS)
	// fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	// fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

    a := 10.00
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := a + b

	fmt.Println("\nC has value", c, "and is of type", reflect.TypeOf(c))
}