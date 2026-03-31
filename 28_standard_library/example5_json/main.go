package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Marshaling (Go struct to JSON)
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	// Unmarshaling (JSON to Go struct)
	var newUser User

	jsonStr := `{"name":"Jane Doe", "age":25, "email":"jane@example.com"}`

	err = json.Unmarshal([]byte(jsonStr), &newUser)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", newUser)
}
