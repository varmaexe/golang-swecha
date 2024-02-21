// StringToInt converts a string to an integer.
// If the conversion fails, it returns an error.
package main

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to integer: %v", err)
	}
	return n, nil
}

func main() {
	// Attempt to convert a valid string to an integer.
	result, err := StringToInt("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Attempt to convert an invalid string to an integer.
	result, err = StringToInt("abc")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
