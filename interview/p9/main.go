package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// printOdd prints odd numbers from 1 to n.
// It listens on oddChan and signals evenChan.
func printOdd(ctx context.Context, n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		select {
		case <-ctx.Done():
			log.Println("printOdd cancelled")
			return
		case <-oddChan:
			fmt.Println(i)
			if i+1 <= n {
				evenChan <- struct{}{}
			}
		}
	}
}

// printEven prints even numbers from 1 to n.
// It listens on evenChan and signals oddChan.
func printEven(ctx context.Context, n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= n; i += 2 {
		select {
		case <-ctx.Done():
			log.Println("printEven cancelled")
			return
		case <-evenChan:
			fmt.Println(i)
			if i+1 <= n {
				oddChan <- struct{}{}
			}
		}
	}
}

// startPrinting starts the goroutines and synchronizes the process.
func startPrinting(n int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	// Start goroutines
	go printOdd(ctx, n, oddChan, evenChan, &wg)
	go printEven(ctx, n, oddChan, evenChan, &wg)

	// Start sequence with odd
	oddChan <- struct{}{}

	// Wait for both to finish or timeout
	wg.Wait()
	log.Println("Finished printing numbers.")
}

func main() {
	startPrinting(10)
}
