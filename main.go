package main

import (
	"example/arrays"
	"example/controlstatements"
	"example/datatypes"
	"example/maps"
	"example/operations"
	"example/slices"
	"fmt"
)

func main() {
	// arrays
	arr := []int{1, 2, 3, 4, 5}
	arrays.PrintArray(arr)

	// slices
	slice := []int{6, 7, 8, 9, 10}
	slices.PrintSlice(slice)

	// dataypes
	p := datatypes.Point{X: 10, Y: 20}
	datatypes.PrintPoint(p)

	rect := datatypes.Rectangle{Width: 5, Height: 10}
	fmt.Println("Area of rectangle:", datatypes.Area(rect))

	// Example usage of operations
	fmt.Println("Addition:", operations.Add(10, 5))
	fmt.Println("Subtraction:", operations.Subtract(10, 5))
	fmt.Println("Multiplication:", operations.Multiply(10, 5))

	result, err := operations.Divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}

	// Example usage of map functions
	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}

	maps.PrintMap(myMap)

	if value, ok := maps.Lookup("banana", myMap); ok {
		fmt.Println("Value of 'banana':", value)
	} else {
		fmt.Println("Key 'banana' not found")
	}

	// if statements example
	controlstatements.BasicIfExample(12)

	// for loops example
	controlstatements.ForLoopExample()

	// switch statements example
	controlstatements.SwitchExample(2)
}
