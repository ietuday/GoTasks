//What Are Pointers?
//A pointer is a variable that stores the memory address of another variable.
//
//In Go:
//
//You use * to declare a pointer or to dereference (access the value).
//
//You use & to get the address of a variable.

package main

import "fmt"

func main() {
	x := 10
	p := &x // p is a pointer to x

	fmt.Println(*p) // prints 10 (dereference)
	*p = 20         // changes the value of x via the pointer
	fmt.Println(x)  // prints 20
}

// Pointers are useful for:
// 1. Modifying variables in functions (pass by reference).
// 2. Avoiding copying large structs or arrays (performance).
// 3. Creating linked data structures (like linked lists, trees, etc.).
// 4. Interfacing with low-level system code or C libraries.
// 5. Implementing certain design patterns (like the Singleton pattern).
// 6. Managing resources (like file handles, network connections, etc.).
// 7. Implementing polymorphism (using interfaces).
// 8. Creating closures (functions that capture variables).
// 9. Implementing concurrency patterns (like channels).
// 10. Creating generic data structures (like maps, slices, etc.).
// 11. Implementing reflection (introspection of types).
// 12. Creating custom data types (like structs, interfaces, etc.).
// 13. Implementing error handling (using the error interface).
//var ptr *int // declares a pointer to an int, but it's nil initially
//ptr = new(int) // allocates zeroed int and gives its pointer
//*ptr = 100
//func update(val *int) {
//    *val = *val + 5
//}
//
//func main() {
//    x := 10
//    update(&x)
//    fmt.Println(x) // 15
//}
//type Person struct {
//    Name string
//}
//
//func main() {
//    p := Person{Name: "Uday"}
//    ptr := &p
//    ptr.Name = "CodeGod"
//    fmt.Println(p.Name) // "CodeGod"
//}
//Go lets you access struct fields through pointers without explicit dereferencing. ptr.Name works the same as (*ptr).Name
// 📦 Why Bother?
// Efficient memory usage (especially in large structs)
//
// Modifying values across scopes
//
// Avoiding unnecessary copying
//
// Shared state across goroutines (careful though, mutexes may be needed)
//
// Symbol	Meaning
// &	Address-of operator
// *	Dereference operator
// *T	Type that’s a pointer to T
// nil	Default zero value of pointer
