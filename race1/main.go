// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/9-if--6ZBn

// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// numbers maintains a set of random numbers.
	numbers []int

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup
	
	mutex sync.Mutex
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// Add a count for each goroutine we will create.
	wg.Add(3)

	// Create three goroutines to generate random numbers.
	go random(10)
	go random(10)
	go random(10)

	// Wait for all the goroutines to finish.
	wg.Wait()

	// Display the set of random numbers.
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}

// random generates random numbers and stores them into a slice.
func random(amount int) {
	// Generate as many random numbers as specified.
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)
		mutex.Lock()
		numbers = append(numbers, n)
		mutex.Unlock()
	}

	// Tell main we are done.
	wg.Done()
}