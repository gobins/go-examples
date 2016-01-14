package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinishes(10, 3, 6, 12, 14, 7, 9)

	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
