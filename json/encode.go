package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family    string
	FirstName string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	name := Name{Family: "Smith", FirstName: "John"}
	email1 := Email{Kind: "Personal", Address: "john@smith.com"}
	email2 := Email{Kind: "Work", Address: "john@smith.com"}
	email := []Email{email1, email2}
	person := Person{Name: name, Email: email}
	outFile, err := os.Create("json.json")
	checkError(err)
	encoder := json.NewEncoder(outFile)
	encoder.Encode(person)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
