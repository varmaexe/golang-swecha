// Pointers:
// - Variables that store memory addresses of other variables.
// - Declared using the * symbol (e.g., var x *int).
// - The zero value of a pointer is nil.
// - Use the & operator to get the address of a variable, and * to dereference a pointer.
// - Pointers are crucial for efficient memory management and passing by reference

package main

import "fmt"

func main() {
	// Declare a variable of type int
	var num int = 42

	// Declare a pointer variable of type *int and assign the address of num to it
	var ptr *int = &num

	// Print the value and address of num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)

	// Print the value and address of the variable pointed to by ptr
	fmt.Println("Value of ptr:", *ptr) // Dereference the pointer to get the value
	fmt.Println("Address stored in ptr:", ptr)

	// Update the value of num through the pointer
	*ptr = 100

	// Print the updated value of num
	fmt.Println("Updated value of num:", num)
}
