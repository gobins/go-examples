package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "Is this real life or is this just fantacy?"

	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	for msg := range ch {
		fmt.Print(msg + " ")
	}

}
