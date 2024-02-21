package main

import (
	"errors"
	"fmt"
	"strconv"
)

// User represents a user entity in the application.
type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// ValidateUser validates the user input fields.
func ValidateUser(name, ageStr, email string) (*User, error) {
	// Validate name
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	// Validate age
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, errors.New("invalid age")
	}
	if age < 0 || age > 100 {
		return nil, errors.New("age must be between 0 and 100")
	}

	// Validate email
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	// Create and return user object if all validations pass
	return &User{
		ID:    1,
		Name:  name,
		Age:   age,
		Email: email,
	}, nil
}

func main() {
	// Example usage
	user, err := ValidateUser("Hithesh", "25", "hithesh@swecha.com")
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	// If validation succeeds, use the validated user object
	fmt.Println("Validated user:", *user)
}
