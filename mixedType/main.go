package main

import "fmt"

func main() {
	mixedSlice := []interface{}{42, "Hello", true, 3.14}

	for _, value := range mixedSlice {
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
}
