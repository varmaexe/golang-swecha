// - Goroutines:
// - Lightweight threads managed by the Go runtime.
// - Created using the go keyword before a function call.
// - Enable concurrent execution without the overhead of traditional threads.
package main

import (
	"fmt"
	"time"
)

// Function to print numbers from 1 to 5
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
	}
}

func main() {
	// Start a new goroutine to execute the printNumbers function concurrently
	go printNumbers()

	// Print "Hello" in the main goroutine
	fmt.Println("Hello")

	// Sleep for 1 second to allow the goroutine to finish execution
	time.Sleep(1 * time.Second)
}
