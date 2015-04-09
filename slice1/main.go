// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/5skpU4iFL5

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of the first and second index
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var intsSlice []int

	// Appends numbers to the slice.
	for i := 0; i < 10; i++ {
		intsSlice = append(intsSlice, i)
	}
	
	// Display each value in the slice.
	for _,i := range intsSlice {
		fmt.Println(i)
	}

	// Declare a slice of strings and populate the slice with names.
	stringsSlice := []string{"foo","bar","baz"}

	// Display each index position and slice value.
	for i,v := range stringsSlice {
		fmt.Printf("stringsSlice[%v] is %v\n", i, v)
	}

	// Take a slice of the first and second index positions of the
	// slice of strings.
	ss := stringsSlice[0:2]

	// Display each index position and slice values for the new slice.
	for i,v := range ss {
		fmt.Printf("ss[%v] is %v\n", i, v)
	}
}