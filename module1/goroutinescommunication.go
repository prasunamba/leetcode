package module1

import (
	"fmt"
	"sync"
)

func Goroutinescomm() {
	// Create channels for communication
	dataChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	// Goroutine 1: Send data and receive the result back
	wg.Add(1)
	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when done
		data := 42
		fmt.Println("Goroutine 1: Sending data:", data)
		dataChan <- data // Send data to Goroutine 2

		// Wait for the result from Goroutine 2
		result := <-resultChan
		fmt.Println("Goroutine 1: Received result:", result)
	}()

	// Goroutine 2: Receive data and send back a result
	wg.Add(1)
	go func() {
		defer wg.Done()    // Decrement the WaitGroup counter when done
		data := <-dataChan // Receive data from Goroutine 1
		fmt.Println("Goroutine 2: Received data:", data)
		result := data * 2   // Process the data (e.g., multiply by 2)
		resultChan <- result // Send the result back
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
