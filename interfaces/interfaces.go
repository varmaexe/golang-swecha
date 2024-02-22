// Interfaces:
// - Abstract types that define behavior rather than data.
// - A set of method signatures forms an interface.
// - Types implicitly implement interfaces if they implement all required methods.
// - Interfaces support polymorphism, allowing different types to be treated uniformly

package main

import (
	"fmt"
	"math"
)

// Define an interface named "Shape"
type Shape interface {
	area() float64
}

// Define a struct type "Circle"
type Circle struct {
	radius float64
}

// Define a struct type "Rectangle"
type Rectangle struct {
	width, height float64
}

// Method to calculate the area of a Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Method to calculate the area of a Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	// Create a Circle object
	circle := Circle{radius: 5}

	// Create a Rectangle object
	rectangle := Rectangle{width: 10, height: 5}

	// Calculate the area of the Circle
	printArea(circle)

	// Calculate the area of the Rectangle
	printArea(rectangle)
}

// Function to print the area of a Shape
func printArea(shape Shape) {
	fmt.Println("Area:", shape.area())
}

// In this example:

//     We define an interface named Shape with a method area(),
// which returns a float64.
//     We define two struct types: Circle and Rectangle.
//     Each struct type implements the area() method defined by the Shape interface.
//     In the main() function, we create instances of Circle and Rectangle.
//     // We call the printArea() function with both Circle and Rectangle objects.
// Since both types implement the Shape interface
// 	(by implementing the area() method), they can be passed to printArea().
