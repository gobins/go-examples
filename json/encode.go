package main

import (
	"encoding/json"
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
	person := Person{Name: name, Email: {email1, email2}}

}
