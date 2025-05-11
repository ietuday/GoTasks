package main

import "fmt"

func logMessage(message interface{}) {
	switch v := message.(type) {
	case string:
		fmt.Printf("Log (string): %s\n", v)
	case int:
		fmt.Printf("Log (int): %d\n", v)
	case error:
		fmt.Printf("Log (error): %s\n", v.Error())
	default:
		fmt.Printf("Log (unknown type): %v\n", v)
	}
}

func main() {
	logMessage("This is a log message.")
	logMessage(404)
	logMessage(fmt.Errorf("An error occurred"))
}
