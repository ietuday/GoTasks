package main

import "fmt"

func main() {
	// Maps in GO
	// Maps are unordered collections of key-value pairs.
	// They are similar to dictionaries in Python or hash tables in other languages.
	// Maps are created using the built-in make function or using a map literal.
	// The key type must be comparable, which means it can be compared using == and != operators.
	// The value type can be any type, including other maps.
	// Maps are reference types, which means that when you assign a map to another variable, both variables refer to the same map.
	// When you modify the map using one variable, the changes are reflected in the other variable as well.
	// Maps are not safe for concurrent use, which means that if multiple goroutines access a map at the same time, it can lead
	// to data
	// Creating a map with string keys and int values
	myMap := make(map[string]int)

	// Adding key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["cherry"] = 15

	// Printing the map
	for key, value := range myMap {
		println(key, value)
	}
	// Accessing a value using its key
	value := myMap["banana"]
	println("Value for key 'banana':", value)
	// Checking if a key exists in the map
	if val, exists := myMap["grape"]; exists {
		println("Value for key 'grape':", val)
	} else {
		println("Key 'grape' does not exist in the map")
	}
	// Deleting a key-value pair from the map
	delete(myMap, "apple")
	println("Map after deleting key 'apple':")
	for key, value := range myMap {
		println(key, value)
	}
	// Iterating over a map
	for key, value := range myMap {
		println("Key:", key, "Value:", value)
	}
	// Creating a nested map
	nestedMap := make(map[string]map[string]int)
	nestedMap["fruits"] = make(map[string]int)
	nestedMap["fruits"]["apple"] = 5
	nestedMap["fruits"]["banana"] = 10
	nestedMap["vegetables"] = make(map[string]int)
	nestedMap["vegetables"]["carrot"] = 3
	nestedMap["vegetables"]["broccoli"] = 7
	// Printing the nested map
	for category, items := range nestedMap {
		println("Category:", category)
		for item, quantity := range items {
			println("Item:", item, "Quantity:", quantity)
		}
	}
	// Creating a map with different key and value types
	complexMap := make(map[int]string)
	complexMap[1] = "one"
	complexMap[2] = "two"
	complexMap[3] = "three"
	// Printing the complex map
	for key, value := range complexMap {
		println("Key:", key, "Value:", value)
	}
	// Creating a map with a custom struct as the value type
	type Person struct {
		Name string
		Age  int
	}
	personMap := make(map[string]Person)
	personMap["john"] = Person{Name: "John Doe", Age: 30}
	personMap["jane"] = Person{Name: "Jane Smith", Age: 25}
	// Printing the person map
	for key, person := range personMap {
		println("Key:", key, "Name:", person.Name, "Age:", person.Age)
	}
	// Creating a map with a custom struct as the key type
	type Point struct {
		X int
		Y int
	}
	pointMap := make(map[Point]string)
	pointMap[Point{1, 2}] = "Point A"
	pointMap[Point{3, 4}] = "Point B"
	// Printing the point map
	for key, value := range pointMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	// Creating a map with a slice as the value type
	sliceMap := make(map[string][]int)
	sliceMap["even"] = []int{2, 4, 6, 8}
	sliceMap["odd"] = []int{1, 3, 5, 7}
	// Printing the slice map
	for key, value := range sliceMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	// Creating a map with a function as the value type
	funcMap := make(map[string]func(int, int) int)
	funcMap["add"] = func(a, b int) int { return a + b }
	funcMap["multiply"] = func(a, b int) int { return a * b }

	// Using the functions in the map
	fmt.Println("Addition:", funcMap["add"](3, 4))
	fmt.Println("Multiplication:", funcMap["multiply"](3, 4))

	// Creating a map with interface{} as the value type
	// This allows storing values of different types in the same map
	interfaceMap := make(map[string]interface{})
	interfaceMap["name"] = "Alice"
	interfaceMap["age"] = 25
	interfaceMap["isStudent"] = true

	// Accessing and type asserting values from the interface map
	for key, value := range interfaceMap {
		switch v := value.(type) {
		case string:
			fmt.Printf("Key: %v, Value (string): %v\n", key, v)
		case int:
			fmt.Printf("Key: %v, Value (int): %v\n", key, v)
		case bool:
			fmt.Printf("Key: %v, Value (bool): %v\n", key, v)
		default:
			fmt.Printf("Key: %v, Value (unknown type): %v\n", key, v)
		}
	}

	// Creating a map with a channel as the value type
	channelMap := make(map[string]chan int)
	channelMap["numbers"] = make(chan int, 2)

	// Sending and receiving values from the channel in the map
	channelMap["numbers"] <- 42
	channelMap["numbers"] <- 84
	close(channelMap["numbers"])

	fmt.Println("Values from channel in map:")
	for val := range channelMap["numbers"] {
		fmt.Println(val)
	}
}
