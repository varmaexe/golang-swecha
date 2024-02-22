// Functions:
// - Defined using the func keyword, followed by the function name and parameters.
// - Parameters can have explicit types, and multiple return values are supported.
// - Named return values simplify function signatures.
// - Variadic functions allow variable-length parameter lists using ....
// - Functions are first-class citizens, supporting closures and anonymous functions

package main

import "fmt"

// Function to add two integers and return the result
func add(x, y int) int {
	return x + y
}

// Function to subtract two integers and return the result
func subtract(x, y int) int {
	return x - y
}

func main() {
	// Calling the add function
	sum := add(3, 5)
	fmt.Println("3 + 5 =", sum)

	// Calling the subtract function
	difference := subtract(10, 3)
	fmt.Println("10 - 3 =", difference)
}
