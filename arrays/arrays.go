// Arrays:
// - Fixed-size collections of elements of the same type.
// - Declared using [size]type syntax.
// - Indexing starts at 0, and elements can be accessed using square brackets.

package arrays

import "fmt"

func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
