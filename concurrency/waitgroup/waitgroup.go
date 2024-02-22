WaitGroups:
- Purpose:
- The sync package in Go provides the WaitGroup type for synchronizing Goroutines.
- It allows one Goroutine to wait for a collection of Goroutines to finish their execution before
proceeding.
- Usage:
import (
"sync"
)
var wg sync.WaitGroup
func main() {
// Launch Goroutines
wg.Add(1)
go myFunction()
// Wait for all Goroutines to finish
wg.Wait()
}
func myFunction() {
defer wg.Done() // Decrement the WaitGroup counter when the Goroutine completes
// Code to be executed concurrently
}

Summary:
- select is a powerful tool for handling multiple channels concurrently, allowing for
non-blocking communication.
- WaitGroups provide a straightforward mechanism for waiting for a collection of Goroutines to
complete their execution, promoting synchronization in concurrent programs.