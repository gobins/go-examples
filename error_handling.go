package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("/Users/Gobi/test/test.py")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}
