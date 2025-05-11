package main

import "fmt"

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	processValue(42)
	processValue("Hello, Go!")
	processValue(true)
	processValue(3.14)
}
