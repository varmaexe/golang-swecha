// Here, the writeToChannel function writes data to the channel (ch). Multiple Goroutines
// concurrently execute this function, safely writing data to the shared channel.
// By leveraging Goroutines and Channels, Go provides a clean and efficient way to handle
// concurrent read and write operations, ensuring data consistency and reducing the likelihood of
// race conditions or other concurrency issues

package main

// Function to write data to the channel
func writeToChannel(ch chan<- int, data int) {
	// Perform some computation on data
	// ...
	ch <- data // Writing to the channel
}

func main() {
	dataChannel := make(chan int)

	// Launch multiple goroutines to write to the channel concurrently
	go writeToChannel(dataChannel, 1)
	go writeToChannel(dataChannel, 2)
	go writeToChannel(dataChannel, 3)

	// Read data from the channel or perform other operations
	// ...
}
