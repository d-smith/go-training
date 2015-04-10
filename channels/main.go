// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/h0nMS_l1rO

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"sync"
	"fmt"
)


// Declare a wait group variable.
var wg sync.WaitGroup

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel used to share the value back and forth.
func goroutine(name string, c chan int) {
	for {
		// Receive on the channel, waiting for the integer. Check
		// if the channel is closed.
		iVal, ok := <-c
		if !ok {
			wg.Done()
			return
		}

		// Display the integer value received.
		fmt.Printf("%s got %d\n", name, iVal)

		// Terminate the goroutine when the value is 10.
		if iVal == 10 {
			// Close the channel.
			close(c)
			// Wait on the wait group to hit 0.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Println("goroutine " + name + " is leaving the building")
			return
		}

		// Increment the value and send in back through
		// the channel.
		c <- iVal + 1
	}
}

// main is the entry point for all Go programs.
func main() {
	// Declare and initialize an unbuffered channel
	// of type int.
	var c = make(chan int)

	// Increment the wait group for each goroutine
	// to be created.
	wg.Add(2)

	// Create the two goroutines.
	go goroutine("thing1",c)
	go goroutine("thing2",c)

	// Send the initial integer value into the channel.
	c <- 1

	// Wait for all the goroutines to finish.
	wg.Wait()
}