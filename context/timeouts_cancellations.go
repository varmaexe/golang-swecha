package main

import (
	"context"
	"fmt"
	"time"
)

// fetchData simulates fetching data from an external service.
// It takes a context and returns the fetched data or an error.
func fetchContent(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "Data fetched successfully", nil
	case <-ctx.Done():
		return "", ctx.Err() // Return context error if context is cancelled
	}
}

func main() {
	// Create a context with a timeout of 1 second
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1() // Ensure cancel1 is called to release resources

	// Perform some operation with ctx1
	result1, err1 := fetchContent(ctx1, 500*time.Millisecond)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Result 1:", result1)
	}

	// Create a context with a timeout of 2 seconds
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2() // Ensure cancel2 is called to release resources

	// Perform some operation with ctx2
	result2, err2 := fetchContent(ctx2, 3*time.Second)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result 2:", result2)
	}

	// Create a context with cancelation
	ctx3, cancel3 := context.WithCancel(context.Background())
	defer cancel3() // Ensure cancel3 is called to release resources

	// Perform some operation with ctx3
	go func() {
		time.Sleep(2 * time.Second)
		cancel3() // Cancel the context after 2 seconds
	}()

	result3, err3 := fetchContent(ctx3, 5*time.Second)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Result 3:", result3)
	}
}

// play with the timelines to understand more efficiently
