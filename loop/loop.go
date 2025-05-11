package main

import "fmt"

func logLoop(name string) {
	fmt.Println("Running:", name)
}
func main() {
	// while loop
	logLoop("while loop")
	i := 1
	for i <= 10 {
		println(i)
		i++
	}

	// infinite loop
	logLoop("infinite loop")
	for {
		println("Infinite loop")
		break // break the infinite loop
	}

	// classic for loop
	logLoop("classic for loop")
	for j := 10; j >= 0; j-- {
		println(j)
	}

	// for range loop
	logLoop("for range loop")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}

	// nested for loop
	logLoop("nested for loop")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			println("i:", i, "j:", j)
		}
	}

	// break and continue
	logLoop("break and continue")
	for k := 1; k <= 5; k++ {
		if k == 3 {
			continue // skip the rest of the loop when k is 3
		}
		if k == 5 {
			break // exit the loop when k is 5
		}
		println(k)
	}

	// labeled break
	logLoop("labeled break")
outerLoop:
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			if m == 2 && n == 2 {
				break outerLoop // exit the outer loop when m is 2 and n is 2
			}
			println("m:", m, "n:", n)
		}
	}

	// labeled continue
	logLoop("labeled continue")
innerLoop:
	for p := 1; p <= 3; p++ {
		for q := 1; q <= 3; q++ {
			if p == 2 && q == 2 {
				continue innerLoop // skip the rest of the inner loop when p is 2 and q is 2
			}
			println("p:", p, "q:", q)
		}
	}

	// for loop with condition
	logLoop("for loop with condition")
	for r := 1; r <= 5; r++ {
		if r%2 == 0 {
			println(r, "is even")
		} else {
			println(r, "is odd")
		}
	}

	// for loop with initialization, condition, and post statement
	logLoop("for loop with initialization, condition, and post statement")
	for s := 1; s <= 5; s++ {
		println("Current value of s:", s)
	}

	// for loop with multiple variables
	logLoop("for loop with multiple variables")
	for t, u := 1, 5; t <= 5 && u >= 1; t, u = t+1, u-1 {
		println("t:", t, "u:", u)
	}

	// for loop with range and map
	logLoop("for loop with range and map")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range ages {
		println(name, "is", age, "years old")
	}

	// for loop with range and string
	logLoop("for loop with range and string")
	str := "Hello"
	for index, char := range str {
		println("Index:", index, "Character:", string(char))
	}

	// for loop with range and array
	logLoop("for loop with range and array")
	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		println("Index:", index, "Value:", value)
	}

	// for loop with range and slice
	logLoop("for loop with range and slice")
	slice := []int{10, 20, 30, 40, 50}
	for index, value := range slice {
		println("Index:", index, "Value:", value)
	}
}

// This code demonstrates various types of loops in Go.
// It includes while loops, infinite loops, classic for loops, for range loops,
// nested loops, break and continue statements, labeled breaks and continues,
// for loops with conditions, initialization, and post statements,
// and for loops with multiple variables.
// The code also includes examples of for loops with range and maps, strings, arrays, and slices.
// The output shows the values of the variables and the results of the loops.
// This code is a simple demonstration of loops in Go.
