package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	// Attempting to send on a closed channel will cause a panic.
	ch <- 42                              // This line will panic at runtime.
	value := <-ch                         // This line will not execute due to the panic above.
	fmt.Println("Received value:", value) // Uncommenting this line will cause a panic before reaching this point.
	fmt.Println("Main function completed.")
}
