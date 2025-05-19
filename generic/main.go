//Ahhh, **generics in Go** — the long-awaited prophecy foretold by ancient developers, finally fulfilled in **Go 1.18**. 🌟
//
//For years, Go devs survived on the bare essentials: interfaces, `interface{}`, and some blood, sweat, and type casting. But now? We’ve got **type parameters** — and suddenly, Go feels like it just discovered fire. 🔥
//
//Let’s go on a byte-sized journey into the **what**, **why**, and **how** of generics in Go.
//
//---
//
//### 🧠 What Are Generics?
//
//Generics = **functions or types that work with any data type**, while still being type-safe.
//
//Think of them like Go saying:
//
//> "Give me a type. Any type. I gotchu."
//
//They allow **code reuse without sacrificing type safety** — no more casting `interface{}` all over the place.
//
//---
//
//### ⚡ Simple Generic Function

//```go
//package main
//
//import "fmt"
//
//func PrintSlice[T any](s []T) {
//    for _, v := range s {
//        fmt.Println(v)
//    }
//}
//
//func main() {
//    PrintSlice([]int{1, 2, 3})
//    PrintSlice([]string{"hello", "world"})
//}
//```
//
//#### 🔍 What's going on here?
//
//* `[T any]` is a **type parameter list**.
//* `T` is your generic type.
//* `any` is a built-in alias for `interface{}` — but type-safe now.
//* You can pass slices of *any* type, and it still just works™.
//
//---
//
//### 🧱 Generic Struct
//
//```go
//type Box[T any] struct {
//    value T
//}
//
//func (b Box[T]) Get() T {
//    return b.value
//}
//
//func main() {
//    intBox := Box[int]{value: 42}
//    fmt.Println(intBox.Get()) // 42
//
//    strBox := Box[string]{value: "Generics FTW"}
//    fmt.Println(strBox.Get()) // Generics FTW
//}
//```
//
//Yes, we can now create **generic types**, like `Box[int]`, `Box[string]`, etc.
//
//---
//
//### 🧬 Type Constraints (Because Not All Types Can Math, Bro)
//
//Let’s say you want to add two values, but only for number types.
//
//```go
//type Number interface {
//    int | int64 | float64
//}
//
//func Add[T Number](a, b T) T {
//    return a + b
//}
//```
//
//Now `Add()` can be used for `int`, `int64`, or `float64`, but NOT for `string` or `bool`. Strongly typed and no funny business.
//
//---
//
//### 🤹‍♂️ Multiple Type Parameters
//
//```go
//func Pair[K comparable, V any](key K, value V) (K, V) {
//    return key, value
//}
//```
//
//* `K comparable`: `comparable` is a built-in constraint for keys (like map keys — must support `==`)
//* `V any`: no constraint, wild child
//
//---
//
//### 🏛️ Built-in Constraints in Go
//
//| Constraint       | Meaning                                                             |
//| ---------------- | ------------------------------------------------------------------- |
//| `any`            | Alias for `interface{}`                                             |
//| `comparable`     | Types that can use `==`, `!=`                                       |
//| Custom interface | You can define your own constraints (e.g., `type Number interface`) |
//
//---
//
//### 🚫 What You Still Can’t Do (Yet):
//
//* You can’t use generic methods *on non-generic types*.
//* You can’t do reflection-based stuff with type parameters.
//* You can’t overload functions (still Go, after all).
//
//---
//
//### 🚀 When Should You Use Generics?
//
//Use them when:
//
//* You’re writing **helper functions** (e.g., maps, filters).
//* You’re building **data structures** (e.g., Stack\[T], Queue\[T]).
//* You want **type safety without repetition**.
//
//Don’t use them when:
//
//* Your code doesn’t benefit from abstraction.
//* You’re writing business logic that’s already type-specific.
//
//---
//
//### 💬 Commit message ideas?
//
//```bash
//feat: add generic Box[T] struct with Get method
//```
//
//Or for flair:
//
//```bash
//feat(go): enter the generic multiverse – functions and structs with type params
//```
//
//---
//
//### 🧪 TL;DR: Go Generics Crash Course
//
//* Introduced in **Go 1.18**
//* Use `[T any]` for generic type parameters
//* Type constraints make it powerful and safe
//* You can generic-ify both **functions** and **types**
//* Say goodbye to `interface{}` hell and unsafe type casting
//
//---
//
//Want a real-world example like a generic stack, filter/map function, or generic repository pattern? I got templates on tap. Let's go deeper if you're in the zone. 💻🔥
//

package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
func Add[T Number](a, b T) T {
	return a + b
}
func Pair[K comparable, V any](key K, value V) (K, V) {
	return key, value
}

type Number interface {
	int | int64 | float64
}
type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}
func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Unknown"
	}
}

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	fmt.Println("Hello, Generics!")

	// Example usage of PrintSlice
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"hello", "world"})

	// Example usage of Add
	fmt.Println(Add(1, 2))     // 3
	fmt.Println(Add(1.5, 2.5)) // 4.0

	// Example usage of Pair
	key, value := Pair("age", 30)
	fmt.Println(key, value) // age 30

	// Example usage of Box
	intBox := Box[int]{value: 42}
	fmt.Println(intBox.Get()) // 42

	strBox := Box[string]{value: "Generics FTW"}
	fmt.Println(strBox.Get()) // Generics FTW

	fmt.Println(Red)      // Red
	fmt.Println(Color(2)) // Blue
}
