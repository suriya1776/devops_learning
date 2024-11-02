package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Separate function that will run as a goroutine
func performTask(taskName string) {
	defer wg.Done() // Decrement the WaitGroup counter when the function completes
	time.Sleep(2 * time.Second)
	fmt.Printf("%s finished\n", taskName)
}

func main() {
	wg.Add(3) // Set counter for 3 goroutines

	// Launch goroutines
	go performTask("Task 1")
	go performTask("Task 2")
	go performTask("Task 3")

	// Wait until all goroutines finish
	wg.Wait()
	fmt.Println("All tasks completed.")
}
