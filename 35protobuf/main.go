package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adasarpan404/protobuf/protobuf"
	"github.com/golang/protobuf/proto"
)

func main() {
	// Create a Person instance and set its values
	person := &protobuf.Person{
		Name: "John Doe",
		Age:  30,
		Hobbies: []string{
			"Coding",
			"Reading",
		},
	}

	// Serialize the person to bytes
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("Error while marshaling:", err)
		return
	}

	// Write the data to a file (optional)
	err = ioutil.WriteFile("person.bin", data, 0644)
	if err != nil {
		fmt.Println("Error while writing to file:", err)
		return
	}

	// Read the data back from the file (optional)
	dataFromFile, err := ioutil.ReadFile("person.bin")
	if err != nil {
		fmt.Println("Error while reading from file:", err)
		return
	}

	// Create an empty Person instance
	newPerson := &protobuf.Person{}

	// Unmarshal the data into the newPerson
	err = proto.Unmarshal(dataFromFile, newPerson)
	if err != nil {
		fmt.Println("Error while unmarshaling:", err)
		return
	}

	// Print the contents of the newPerson
	fmt.Println("Name:", newPerson.GetName())
	fmt.Println("Age:", newPerson.GetAge())
	fmt.Println("Hobbies:", newPerson.GetHobbies())
}
