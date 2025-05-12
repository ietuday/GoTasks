package main

import "fmt"

// Define the interface and structs up front
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

// Implement Speak() for Dog
func (d Dog) Speak() string {
	return "Woof"
}

// Implement Speak() for Cat
func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	// Example usage of the range function
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example usage of the range function with a map
	days := map[string]int{"Monday": 1, "Tuesday": 2, "Wednesday": 3}
	for day, num := range days {
		fmt.Println(day, num)
	}

	// Example usage of the range function with a string
	str := "Hello"
	for i, char := range str {
		fmt.Println(i, string(char))
	}

	// Example usage of the range function with a channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}

	// Example usage of the range function with a slice of structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}

	// Example usage of the range function with a slice of pointers
	type Point struct {
		X int
		Y int
	}
	points := []*Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	for _, point := range points {
		fmt.Println(point.X, point.Y)
	}

	// Example usage of the range function with a slice of interfaces
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
