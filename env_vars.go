package main

import (
    "fmt"
    "os"
)

func main() {
    name := os.Getenv("USER")

    fmt.Println("\nHi,",name)

    fmt.Println(os.Environ())
}
