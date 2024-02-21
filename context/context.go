// In Go, the context package provides a way to manage the lifecycle and cancellation of operations within your program.
//  It's commonly used to propagate deadlines, cancellation signals, and request-scoped values across API boundaries and
// between goroutines.

package main

import (
	"context"
	"fmt"
	"time"
)

// fetchData simulates fetching data from an external service.
// It takes a context and returns the fetched data or an error.
func fetchData(ctx context.Context) (string, error) {
	// Simulate a long-running operation
	select {
	case <-time.After(2 * time.Second):
		return "Data fetched successfully", nil
	case <-ctx.Done():
		return "", ctx.Err() // Return context error if context is cancelled
	}
}

func main() {
	// Create a context with timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Ensure cancel is called to release resources

	// Perform some operation with the context
	result, err := fetchData(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Result:", result)
}

// play with timeouts, and know its use case
