package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxRetries := 5
	retryInterval := 1 // seconds

	for retry := 0; retry <= maxRetries; retry++ {
		fmt.Printf("Attempt %d\n", retry)

		if err := performOperation(); err != nil {
			fmt.Printf("Error: %v\n", err)
			if retry == maxRetries {
				fmt.Println("Max retries reached. Giving up.")
				break
			}
			fmt.Printf("Retrying in %d seconds...\n", retryInterval)
			time.Sleep(time.Duration(retryInterval) * time.Second)
			retryInterval *= 2 // Exponential backoff: double the interval
		} else {
			fmt.Println("Operation successful")
			break
		}
	}
}

func performOperation() error {
	// Simulate an operation that may fail randomly
	if rand.Intn(10) < 8 {
		return fmt.Errorf("temporary error")
	}
	return nil
}
