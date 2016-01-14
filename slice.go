package main

import (
	"fmt"
)

func main() {

	myCourses := make([]string, 5, 20)
	myChoices := []string{"Docker", "Vagrant", "Packer"}

	fmt.Printf("Length is: %d.\nCapacity is: %d\n", len(myCourses), cap(myCourses))
	fmt.Printf("\nLength is: %d.\nCapacity is: %d\n", len(myChoices), cap(myChoices))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlice[3] = 0

	fmt.Println(mySlice)

	sliceOfSlice := mySlice[3:]

	fmt.Println(sliceOfSlice)

}
