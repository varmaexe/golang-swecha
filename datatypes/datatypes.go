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
