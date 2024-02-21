// Maps:
// - Key-value pairs for efficient data retrieval.
// - Declared using map[keyType]valueType.
// - Operations include adding, updating, deleting elements, and checking existence.

package maps

import "fmt"

// Define functions to work with maps here

func PrintMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func Lookup(key string, m map[string]int) (int, bool) {
	value, ok := m[key]
	return value, ok
}
