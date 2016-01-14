package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v\n", key, value)
	}

	fmt.Println(testMap)
	testMap["F"] = 1980

	fmt.Println(testMap)

	delete(testMap, "A")
	fmt.Println(testMap)
}
