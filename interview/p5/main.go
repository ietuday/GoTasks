package main

import (
	"fmt"
)

func main() {
	var ch chan int

	go func() {
		ch = make(chan int)
		ch <- 42
	}()
	value := <-ch
	fmt.Println("Received value:", value)
}
