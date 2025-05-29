package main

import (
	"fmt"
	"sync"
)

// printOdd prints odd numbers up to `n` using synchronization via channels.
// It waits on oddChan and signals evenChan after each odd number is printed.
func printOdd(n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it returns
	for i := 1; i <= n; i += 2 {
		<-oddChan // Wait for signal from even goroutine to print an odd number
		fmt.Println(i)

		// If more numbers are left, signal even goroutine to continue
		if i+1 <= n {
			evenChan <- struct{}{}
		}
	}
}

// printEven prints even numbers up to `n` using synchronization via channels.
// It waits on evenChan and signals oddChan after each even number is printed.
func printEven(n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it returns
	for i := 2; i <= n; i += 2 {
		<-evenChan // Wait for signal from odd goroutine to print an even number
		fmt.Println(i)

		// If more numbers are left, signal odd goroutine to continue
		if i+1 <= n {
			oddChan <- struct{}{}
		}
	}
}

func main() {
	n := 10 // Maximum number to print

	// Create unbuffered channels for synchronization
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2) // We're launching two goroutines

	// Launch goroutines to print odd and even numbers
	go printOdd(n, oddChan, evenChan, &wg)
	go printEven(n, oddChan, evenChan, &wg)

	// Kick off the printing by signaling the odd goroutine first
	oddChan <- struct{}{}

	// Wait for both goroutines to finish
	wg.Wait()

	fmt.Println("All numbers printed.")
}
