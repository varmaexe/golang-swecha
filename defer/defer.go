// Defer:
// - Used to schedule a function call to be run after the function containing the defer statement
// completes.
// - Commonly used for cleanup operations.
// - Deferred functions are executed in Last In, First Out (LIFO) orde

package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Deferred functions are executed in Last In First Out order
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Deferred function calls scheduled")
}
