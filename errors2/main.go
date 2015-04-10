// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/x6UimVQMMQ

// Create a custom error type called appError that contains three fields, Err error,
// Message string and Code int. Implement the error interface providing your own message
// using these three fields. Write a function called checkFlag that accepts a bool value.
// If the value is false, return a pointer of your custom error type initialized as you like.
// If the value is true, return a default error. Write a main function to call the
// checkFlag function and check the error for the concrete type.
package main

// Add imports.
import (
	"errors"
	"fmt"
)

// Declare a struct type named appError with three fields, Err of type error,
// Message of type string and Code of type int.
type appError struct {
	Err     error
	Message string
	Code    int
}

// Declare a method for the appError struct type that implements the
// error interface.
func (e *appError) Error() string {
	return fmt.Sprintf("%s: %d - %s", e.Err, e.Code, e.Message)
}

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
func checkFlag(customError bool) error {
	// If the parameter is false return an appError.
	if customError {
		theError := errors.New("boom")
		return &appError{theError, "Made an error", 1}
	}

	// Return a default error.
	return errors.New("standard error")
}

// main is the entry point for the application.
func main() {
	// Call the checkFlag function to simulate an error of the
	// concrete type.
	err := checkFlag(true)

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {
	// Apply the case for the default error type.
	case *appError:
		fmt.Println(e)

	// Apply the default case.
	default:
		fmt.Println(e)
	}
}
