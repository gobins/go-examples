package main

import "fmt"

func main() {

	ch := make(chan string, 1)
	ch <- "Hello Go"
	fmt.Println(<-ch)

}
