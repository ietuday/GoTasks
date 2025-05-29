package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s!", name)
	return message
}
