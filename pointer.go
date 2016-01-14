package main

import (
	"fmt"
)

func main() {
	name := "Gobin"
	course := "Go Fundamentals"

	fmt.Println("\nHi", name, "you're currently wathcing", course)

	changeCourse(&course)

	fmt.Println("\nYour new course is", course)
}

func changeCourse(course *string) string {
	*course = "Go Dive"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}
