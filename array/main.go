package main

import "fmt"

// main is the entry point of the program that demonstrates the usage of arrays and slices in Go.
// It performs the following operations:
// - Creates and initializes an array, prints its elements, length, and capacity.
// - Calculates and prints the sum of the elements in the array.
// - Creates and initializes a slice, prints its elements, length, and capacity.
// - Demonstrates appending elements to a slice and prints the updated slice.
// - Demonstrates deleting an element from a slice and prints the updated slice.
// - Demonstrates inserting an element into a slice and prints the updated slice.
// It includes examples of array initialization, accessing elements, calculating the sum of elements,
// and performing operations on slices such as appending, deleting, and inserting elements.
func main() {
	// array Demotration
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println("Array length:", len(arr))
	fmt.Println("Array capacity:", cap(arr))
	var sum [5]int
	sum[0] = 1
	sum[1] = 2
	sum[2] = 3
	sum[3] = 4
	sum[4] = 5
	println("Sum of array:", sum[0]+sum[1]+sum[2]+sum[3]+sum[4])

	// slice Demotration
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))
	slice = append(slice, 6)
	fmt.Println("Slice after append:", slice)
	slice = append(slice, 7, 8, 9)
	fmt.Println("Slice after append:", slice)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("Slice after delete:", slice)
	slice = append(slice[:2], 10)
	fmt.Println("Slice after insert:", slice)
	slice = append(slice[:2], 11)
	fmt.Println("Slice after insert:", slice)
	slice = append(slice[:2], 12)
	fmt.Println("Slice after insert:", slice)

	// Additional examples of arrays
	arr2 := [3]string{"Go", "Python", "Java"}
	fmt.Println("String Array:", arr2)
	fmt.Println("First element of array:", arr2[0])

	// Modifying an array
	arr2[1] = "C++"
	fmt.Println("Modified Array:", arr2)

	// Additional examples of slices
	slice2 := []float64{1.1, 2.2, 3.3}
	fmt.Println("Float Slice:", slice2)

	// Slicing a slice
	subSlice := slice2[1:]
	fmt.Println("Sub-slice:", subSlice)

	// Copying a slice
	copySlice := make([]float64, len(slice2))
	copy(copySlice, slice2)
	fmt.Println("Copied Slice:", copySlice)

	// Extending a slice
	extendedSlice := append(slice2, 4.4, 5.5)
	fmt.Println("Extended Slice:", extendedSlice)

	// Creating a slice with make
	madeSlice := make([]int, 5, 10)
	fmt.Println("Slice created with make:", madeSlice)
	fmt.Println("Length of madeSlice:", len(madeSlice))
	fmt.Println("Capacity of madeSlice:", cap(madeSlice))
}
