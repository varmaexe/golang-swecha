// Sqrt returns the square root of a number.
// If the number is negative, it returns an error.
package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return math.Sqrt(x), nil
}

func main() {
	// Calculate the square root of a positive number.
	result, err := Sqrt(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}

	// Attempt to calculate the square root of a negative number.
	result, err = Sqrt(-16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}
}
