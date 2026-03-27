package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("==========BASIC============")

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Println("==========Example 1============")

	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	// Create a slice of user values with two users.
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of users switching
	// semantics. Not Good!
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Exception example: Using pointer semantics
	// for a collection of strings.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}

	fmt.Println("==========Example 2============")

	// Declare a variable of type duration set to
	// its zero value.
	var dur duration

	// Change the value of dur to equal
	// five hours.
	dur.setHours(5)

	// Display the new value of dur.
	fmt.Println("Hours:", dur.hours())

	fmt.Println("==========Example Celsius Language IntSlice============")

	temp := Celsius(100)
	fmt.Printf("%.1f°C = %.1f°F\n", temp, temp.ToFahrenheit()) // 100.0°C = 212.0°F

	lang := Language("go")
	fmt.Println(lang.IsValid()) // → true
	fmt.Println(lang.Upper())   // → GO

	nums := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println(nums.Sum()) // → 31
	fmt.Println(nums.Max()) // → 9
}
