package main

func main() {
	// while loop
	i := 1
	for i <= 10 {
		println(i)
		i++
	}

	// infinite loop
	for {
		println("Infinite loop")
		break // break the infinite loop
	}
	// classic for loop
	for j := 10; j >= 0; j-- {
		println(j)
	}
}
