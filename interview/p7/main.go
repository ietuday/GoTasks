package main

import (
	"fmt"
)

func printOdd(oddCh, evenCh chan int, n int) {
	for i := 1; i <= n; i++ {
		<-oddCh // Wait for the signal to print
		fmt.Println(i)
		evenCh <- 1 // Signal to print the next even number
	}
}

func printEven(oddCh, evenCh chan int, n int) {
	for i := 1; i <= n; i++ {
		<-evenCh // Wait for the signal to print
		fmt.Println(i)
		oddCh <- 1 // Signal to print the next even number
	}
}

func main() {
	n := 10
	// Create channels for odd and even numbers
	oddCh := make(chan int)
	evenCh := make(chan int)

	// Start goroutines for printing odd and even numbers
	go printOdd(oddCh, evenCh, n)
	go printEven(oddCh, evenCh, n)

	oddCh <- 1 // Start the odd sequence

	select {} // Block forever to keep the main goroutine alive

}
