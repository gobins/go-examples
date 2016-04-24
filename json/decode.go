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
	Family   string
	FirsName string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	var person Person
	inFile, err := os.Open("json.json")
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err1 := decoder.Decode(&person)
	checkError(err1)
	inFile.Close()

	fmt.Println(person.Name.FirsName)
	//fmt.Println(person.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}

// func (p Person) String() string {
// 	s := p.Name.FirsName + "" + p.Name.Family
// 	for _, v := range p.Email {
// 		s += "\n" + v.Kind + ":" + v.Address
// 	}
// 	return s
// }
