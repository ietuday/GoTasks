package main

import (
	"fmt"
)

func send(ch chan int) {
	ch <- 42
}

func main() {
	ch := make(chan int)
	go send(ch)
	value := <-ch
	fmt.Println("Received value:", value)
	fmt.Println("Main function completed.")
}
