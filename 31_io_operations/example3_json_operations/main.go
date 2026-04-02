package main

import (
	"encoding/json"
	"fmt"
)

// Struct with tags
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {

	/*
		Key Concepts
			omitempty → skip empty fields
			Tags control JSON field names
			Always pass pointer in Unmarshal
	*/
	user := User{
		Name: "Aditya",
		Age:  30,
	}

	// Marshal (struct → JSON)
	jsonData, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	// Unmarshal (JSON → struct)
	var newUser User

	// Assign email value
	newUser.Email = "a@g.com"

	// Always pass pointer in Unmarshal
	err = json.Unmarshal(jsonData, &newUser)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", newUser)
}
