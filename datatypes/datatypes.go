// Data Types:
// - Golang supports standard data types like int, float, bool, string.
// - Additional data types include complex numbers, runes, and pointers.
// - The var keyword is used for variable declaration, and Go infers the data type.

package datatypes

import "fmt"

// Define your custom data types here

type Point struct {
	X, Y int
}

type Rectangle struct {
	Width, Height int
}

// Example function to demonstrate usage of custom data types
func PrintPoint(p Point) {
	fmt.Printf("Point: (%d, %d)\n", p.X, p.Y)
}

func Area(rect Rectangle) int {
	return rect.Width * rect.Height
}
