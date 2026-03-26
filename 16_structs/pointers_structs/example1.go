package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Value receiver → copy of struct (safe, no mutation)
func (p Person) Greet() string {
	return "Hello, " + p.Name
}

// Pointer receiver → works on original, allows mutation
func (p *Person) HaveBirthday() {
	p.Age++ // mutates the original!
}

func main() {
	person := Person{Name: "Charlie", Age: 28}
	fmt.Println(person)      // value receiver works on value
	(&person).HaveBirthday() // or just:
	fmt.Println(person)      // Go auto-takes address for pointer methods!

	var person2 Person
	person2 = person
	(&person2).HaveBirthday()
	fmt.Println(person)
	fmt.Println(person2)

	fmt.Println("====================================================")

	// Creating with pointers
	email := "alice@example.com"
	settings := UserSettings{Theme: "dark", Notifications: true}

	u := User{
		Name:  "Alice",
		Email: &email,
		Posts: []*Post{
			{Title: "Go Pointers", Content: "Learn about pointers in Go"},
			{Title: "Go Pointers 2", Content: "Deep dive into structs"},
			{Title: "Go Pointers 3", Content: "Best practices with pointers"},
		},
		Settings: &settings,
	}

	fmt.Printf("%+v\n", u) // This may still show addresses
	fmt.Println(u)         // This will now use the clean String() method above
	fmt.Println(u.Posts)
}
