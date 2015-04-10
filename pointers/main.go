// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/9_MSdcdlNQ

// Declare a struct type named Point with two fields, X and Y of type int.
// Implement a factory function for this type and a method that accept a value
// of this type and calculates the distance between the two points. What is
// the nature of this type?
package main

// Add imports.
import (
	"fmt"
	"math"
)

// Declare struct type named Point that represents a point in space.
type Point struct {
	X int
	Y int
}

// Declare a function named New that returns a Point based on X and Y
// positions on a graph.
func New(x int, y int) Point {
	return Point{X: x, Y: y}
}

// Declare a method named Distance that finds the length of the hypotenuse
// between two points. Pass one point in and return the answer.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
// Use the math.Pow and math.Sqrt functions.
func (thisPoint Point) Distance(thatPoint Point) float64 {
	xDiff := float64(thisPoint.X - thatPoint.X)
	yDiff := float64(thisPoint.Y - thatPoint.Y)

	return math.Sqrt(math.Pow(xDiff, 2.0) + math.Pow(yDiff, 2.0))
}

// main is the entry point for the application.
func main() {
	// Declare the first point.
	p1 := New(1, 1)

	// Declare the second point.
	p2 := New(2, 1)

	// Calculate the distance and display the result.
	fmt.Println(p1, p2)
	fmt.Println(p1.Distance(p2))
}
