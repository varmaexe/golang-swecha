package controlstatements

import "fmt"

// if statements
func BasicIfExample(x int) {
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}
}

// for loops
func ForLoopExample() {
	// Standard for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Range-based loop
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// Infinite loop
	// for {
	//     fmt.Println("This is an infinite loop")
	//     // Uncomment the above line to run the infinite loop (Ctrl+C to stop)
	// }
}

// switch statements
func SwitchExample(x int) {
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is neither 1 nor 2")
	}
}
