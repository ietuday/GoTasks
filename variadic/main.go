package main

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Variadic function
	// A variadic function is a function that can take a variable number of arguments.
	// In Go, you can define a variadic function by using the ellipsis (...) before the type of the last parameter.
	// For example, func sum(nums ...int) int { ... } is a variadic function that takes a variable number of int arguments.

	// Example of a variadic function
	// sum(1, 2, 3, 4, 5)
	// The sum function takes a variable number of int arguments and returns their sum.
	// The function can be called with any number of int arguments, including zero arguments.
	// The function can also be called with a slice of int arguments by using the ... operator.
	// For example, sum([]int{1, 2, 3}...) is equivalent to sum(1, 2, 3).
	// The ... operator is used to unpack the slice into individual arguments.
	// Define the variadic function

	// Call the variadic function and print the result
	result := sum(1, 2, 3, 4, 5)
	println("The sum is:", result)
	// Call the variadic function with a slice of int arguments
	slice := []int{1, 2, 3}
	result = sum(slice...)
	println("The sum of the slice is:", result)
	// Call the variadic function with no arguments
	result = sum()
	println("The sum of no arguments is:", result)
	// Call the variadic function with a single argument
	result = sum(10)
	println("The sum of a single argument is:", result)
	// Call the variadic function with a slice of int arguments and a single argument
	result = sum(slice[0], slice[1], slice[2], 10)
	println("The sum of the slice and a single argument is:", result)
	// Call the variadic function with a slice of int arguments and no arguments
	result = sum(slice[0], slice[1], slice[2])
	println("The sum of the slice and no arguments is:", result)
}
