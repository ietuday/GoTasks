package main

import "fmt"

func process(data interface{}) {
	switch v := data.(type) {
	case []int:
		fmt.Println("Processing a slice of integers:", v)
	case map[string]string:
		fmt.Println("Processing a map of strings:", v)
	case string:
		fmt.Println("Processing a string:", v)
	default:
		fmt.Printf("Unknown data type: %T\n", v)
	}
}

func main() {
	process([]int{1, 2, 3})
	process(map[string]string{"key": "value"})
	process("Hello, Go!")
	process(42)
}
