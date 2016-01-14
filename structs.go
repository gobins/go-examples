package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	DockerSwarm := courseMeta{
		Author: "Gobin Sougrakpam",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerSwarm)
	fmt.Printf("\nThe author is %s and level %s", DockerSwarm.Author, DockerSwarm.Level)

}
