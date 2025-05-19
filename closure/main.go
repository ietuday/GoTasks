// A closure in Go is a function value that references variables from outside its body. It can access and modify these variables even after the outer function has returned.

// Basically:

// "A closure closes over its environment."

package main

import "fmt"

func main() {
	counter := createCounter()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// In this example, the `createCounter` function returns a closure that increments and returns the value of `count` each time it is called. The `count` variable is captured by the closure, allowing it to maintain its state between calls.
// This is a powerful feature of Go, as it allows you to create functions with state without using global variables or complex data structures.
// Closures are often used in Go for callbacks, event handlers, and other scenarios where you need to maintain state across function calls without using global variables.
// Closures can also be used to create function factories, where you can create functions with specific behavior based on the parameters passed to the outer function.
// Here's another example of using closures to create a function factory:
//func createMultiplier(factor int) func(int) int {
//	return func(x int) int {
//		return x * factor
//	}
//}
//func main1() {
//	double := createMultiplier(2)
//	triple := createMultiplier(3)

//	fmt.Println(double(5)) // 10
//	fmt.Println(triple(5)) // 15
//}
//Closures can be used for:

//Memoization (caching computed results)

//Encapsulation (like private variables)

//Functional-style programming (partial application, currying, etc.)

//Generating factory functions
