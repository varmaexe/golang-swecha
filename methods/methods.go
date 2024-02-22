// Methods:
// - Functions associated with a type, enabling object-oriented programming.
// - Receiver (type on which the method operates) specified between the func keyword and the
// method name.
// - Methods can be attached to both user-defined and built-in types.

package main

import "fmt"

// Define a struct type named "Rectangle"
type Rectangle struct {
	width, height float64
}

// Method area() calculates and returns the area of a Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Method perimeter() calculates and returns the perimeter of a Rectangle
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	// Create a new Rectangle object
	rect := Rectangle{width: 10, height: 5}

	// Call the area() method on the Rectangle object
	fmt.Println("Area of the rectangle:", rect.area())

	// Call the perimeter() method on the Rectangle object
	fmt.Println("Perimeter of the rectangle:", rect.perimeter())
}
