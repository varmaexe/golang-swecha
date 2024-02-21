// how to use values in context
package main

import (
	"context"
	"fmt"
)

type userIDKey struct{}

func main() {
	ctx := context.Background()

	// Add a value to the context with a custom key type
	ctx = context.WithValue(ctx, userIDKey{}, 123)

	// Retrieve the value from the context
	userID, ok := ctx.Value(userIDKey{}).(int)
	if !ok {
		fmt.Println("Error: userID not found in context")
		return
	}

	// Use the retrieved value
	fmt.Printf("User ID: %d\n", userID)

	// Pass the context to a function
	processUser(ctx, userID)
}

func processUser(ctx context.Context, userID int) {
	fmt.Printf("Processing user with ID: %d\n", userID)

	// Access the value from the context within the function, note that we are type asserting here to ensure the value we get is int
	userID2, ok := ctx.Value(userIDKey{}).(int)
	if !ok {
		fmt.Println("Error: userID not found in context (within function)")
		return
	}

	fmt.Printf("User ID within function: %d\n", userID2)
}
