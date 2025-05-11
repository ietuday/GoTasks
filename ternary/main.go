package main

func main() {
	// implement ternary operator
	// Go does not have a built-in ternary operator like some other languages (e.g., C, Java, JavaScript).
	// However, you can achieve similar functionality using a simple if-else statement or a function.

	// Example using if-else statement
	a := 10
	b := 20
	result := 0
	if a > b {
		result = a
	} else {
		result = b
	}
	println(result) // Output: 20

	// Example using a function
	result = ternary(a > b, a, b)
}

// ternary function defined outside main
func ternary(condition bool, trueValue, falseValue int) int {
	if condition {
		return trueValue
	}
	return falseValue
}
