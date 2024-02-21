// Slices:
// - Dynamic and flexible alternatives to arrays.
// - Created using the make function or slicing existing arrays/slices.
// - Support appending and slicing operations.

package slices

import "fmt"

func PrintSlice(slice []int) {
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
