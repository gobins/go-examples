package main

import (
	"fmt"
)

func main() {

	courseInProg := []string{"Docker", "Clustering", "Kubernetes"}
	courseCompleted := []string{"Docker"}

	for _, i := range courseInProg {
		fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("Course", j, "is already completed")
			}
		}
	}
}
