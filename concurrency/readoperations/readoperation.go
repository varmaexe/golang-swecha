// n the example above, the readFromChannel function continuously reads data from the channel
// (ch). Multiple Goroutines can concurrently write to the channel, and the reading Goroutine
// safely processes the data

package main

import "fmt"

// Function to continuously read from the channel
func readFromChannel(ch <-chan int) {
	for {
		data := <-ch // Reading from the channel
		// Process the data
		fmt.Println("Received data:", data)
	}
}

func main() {
	dataChannel := make(chan int)

	// Start a goroutine to continuously read from the channel
	go readFromChannel(dataChannel)

	// Send data to the channel from other goroutines
	dataChannel <- 42
	// ...
}
