// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/HPJvLJoupp

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import (
	"fmt"
)

// Add user type and provide comment.
type user struct {
	name string
	email string
	age int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a composite literal.
	user1 := user {
		name: "User 1",
		email: "user1@foo.om",
		age: 42,
	} 

	// Display the field values.
	fmt.Println(user1)
	

	// Declare a variable using an anonymous struct.
	anotherUser := struct {
		userName string
		anEmail string
		age int
	}{
		userName: "foo",
		anEmail: "anon@dev.null",
		age: 42,
	}

	// Display the field values.
	fmt.Println(anotherUser)
}
