// Single return languages in c
// Throwing exceptions in C++
// Checked exceptions in Java
// Returning exceptions in Go

package main

import (
	"errors"
	"fmt"
)

// Divide function divides two integers.
// It returns the result of the division and an error if the denominator is zero.
func Divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		// Return an error if the denominator is zero.
		return 0, errors.New("cannot divide by zero")
	}
	// Perform the division if the denominator is non-zero.
	return numerator / denominator, nil
}

func main() {
	// Perform a division.
	result, err := Divide(10, 5)
	if err != nil {
		// Handle the error if it occurred.
		fmt.Println("Error:", err)
	} else {
		// Print the result if no error occurred.
		fmt.Println("Result:", result)
	}

	// Attempt a division by zero.
	result, err = Divide(10, 0)
	if err != nil {
		// Handle the error if it occurred.
		fmt.Println("Error:", err)
	} else {
		// Print the result if no error occurred.
		fmt.Println("Result:", result)
	}
}
