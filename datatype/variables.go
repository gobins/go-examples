package main

import (
    "fmt"
    "reflect"
)

var (
	name = "Gobin"
	course = "Golang"
	module = "Go Fundamentals"
)

func main() {

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

}