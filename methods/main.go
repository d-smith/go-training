// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import (
	"fmt"
)

// Declare a struct that represents a ball player.
type Baller struct {
	name   string
	atBats float64
	hits   float64
}

// Include field called name, atBats and hits.

// Declare a method that calculates the batting average for a batter.
func (b *Baller) average() float64 {
	return b.hits / b.atBats
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	players := []Baller{{name: "a", atBats: 1000.0, hits: 10.0}, {name: "b", atBats: 1000.0, hits: 354.0}}

	// Display the batting average for each player in the slice.
	for _, baller := range players {
		fmt.Printf("%s hit %v\n", baller.name, baller.average())
	}
}
