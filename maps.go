package main

import (
	"fmt"
)

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["West Coast Eagles"] = 6
	leagueTitles["Freo Dockers"] = 4

	recentHead2Head := map[string]int{
		"West Coast Eagles": 5,
		"Freo Dockers":      0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)

}
