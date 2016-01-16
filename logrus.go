package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "Dog",
	}).Info("A Dog appears")
	log.Error("This is an error")
	log.SetLevel(log.DebugLevel)
	log.Fatal("Oh my Gerd")
}
