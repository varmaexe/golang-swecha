Select and WaitGroups
Select Statement:
- Purpose:
- The select statement in Go is used to handle multiple channels concurrently.
- It allows Goroutines to wait on multiple communication operations, making it useful for
synchronization in concurrent programs.
- Syntax:
go
select {
case <-ch1:
// Code to execute when data is received from ch1
case ch2 <- data:
// Code to execute when data is sent to ch2
default:
// Code to execute when no communication is ready
}