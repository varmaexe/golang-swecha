// Channels:
// - Communication mechanism between Goroutines.
// - Created using the make function: ch := make(chan Type).
// - Send data into a channel using ch <- data, and receive using data := <-ch.
// - Supports synchronization and data sharing between Goroutines.

package main

import (
	"fmt"
)

// Function to send numbers into a channel
func sendNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send value into the channel
	}
	close(ch) // Close the channel after sending all values
}

// Function to receive numbers from a channel
func receiveNumbers(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	// Create an unbuffered channel of type int
	ch := make(chan int)

	// Start a goroutine to send numbers into the channel
	go sendNumbers(ch)

	// Start another goroutine to receive numbers from the channel
	go receiveNumbers(ch)

	// Sleep for a moment to allow goroutines to complete
	// (In a real-world scenario, you would use synchronization techniques like wait groups)
	// This sleep is added for demonstration purposes only.
	// In practice, you would ensure proper synchronization to wait for goroutines to complete.
	//time.Sleep(1 * time.Second)

	// Waiting for user input to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exiting...")
}
