package main

func main() {
	// Example usage of function
	// Define a function that takes a slice of integers and returns their sum
	// Define a function that takes a slice of integers and returns their sum
	// 		{"Bob", 25},
	// 	}
	// 	{"Charlie", 35},
	// Function to calculate the sum of a slice of integers
	sum := func(numbers []int) int {
		total := 0
		for _, num := range numbers {
			total += num
		}
		return total
	}

	// Example usage
	numbers := []int{1, 2, 3, 4, 5}
	result := sum(numbers)
	println("The sum is:", result)
}
