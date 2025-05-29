// Reverse a string in Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Reversed string:", reverse("Hello, World!"))
	fmt.Println("Reversed string:", reverse("Go is awesome!"))
}

func reverse(s string) string {
	// Convert the string to a slice of runes to handle multi-byte characters correctly
	// This is necessary because Go strings are UTF-8 encoded, and a rune is an alias for int32
	// which can represent any Unicode code point.
	runes := []rune(s)
	// Reverse the slice of runes in place
	// This is done by swapping the elements from the start and end towards the center.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert the slice of runes back to a string and return it
	return string(runes)
}
