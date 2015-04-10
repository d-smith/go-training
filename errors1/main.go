// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rt3O-7ndtJ

// Create two error variables, one called InvalidValueError and the other
// called AmountToLargeError. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the InvalidValueError. Check the value for greater than $1,000 and if it is,
// reutrn the AmountToLargeError. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

// Add imports.
import (
	"errors"
	"fmt"
)

// Declare an error variable named InvalidValueError using the New
// function from the errors package.
var InvalidValueError = errors.New("Invalid Value")

// Declare an error variable named AmountToLargeError using the New
// function from the errors package.
var AmountTooLargeError = errors.New("Amount too large")

// Declare a function named checkAmount that accepts a value of
// type float64 and returns an error interface value.
func checkAmount(checkAmount float64) error {
	// Is the parameter equal to zero. If so then return
	// the error variable.
	if checkAmount == 0 {
		return InvalidValueError
	}

	// Is the parameter greater than 1000. If so then return
	// the other error variable.
	if checkAmount > 1000 {
		return AmountTooLargeError
	}

	// Return nil for the error value.
	return nil
}

// main is the entry point for the application.
func main() {
	// Call the checkAmount function and check the error. Then
	// use a switch/case to compare the error with each variable.
	// Add a default case. Return if there is an error.
	if err := checkAmount(10); err != nil {
		switch err {
		case InvalidValueError:
			fmt.Println("Got an error sayin amount value is invalid")
			return
		case AmountTooLargeError:
			fmt.Println("Error says the amount was too large")
			return
		}
	}

	// Display everything is good.
	fmt.Println("All is in order")
}
