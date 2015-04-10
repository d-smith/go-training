// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/kxcitARJZH

// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

// Add Imports.
import (
	"sync"
	"fmt"
)

// Declare a constant and set the value for the number of goroutines.
const noGoroutines = 20

// Declare a constant and set the value for the number of buffers.
const noBuffers = 4

// Declare a wait group.
var wg sync.WaitGroup

// Declare a buffered channel of type string and initialize it.
var bc = make(chan string, noBuffers)

// Declare a function for the goroutine that will process work
// from the buffered channel.
func worker(c chan string ) {
	// Receive a string from the channel.
	s := <- c
	
	// Display the string.
	fmt.Println("got " + s)

	// Send the string back into the channel.
	c <-s

	// Tell main this goroutine is done.
	wg.Done()
}

// main is the entry point for all Go programs.
func main() {
	// Increment the wait group
	wg.Add(noGoroutines)

	// Create the number if goroutines based on the
	// value of the constant.
	for i := 0; i < noGoroutines; i++ {
		go worker(bc)
	}

	// Send resources into the buffered channel.
	bc <- "A"
	bc <- "B"
	bc <- "C"
	bc <- "D"

	// Wait for all the work to get done.
	wg.Wait()
}