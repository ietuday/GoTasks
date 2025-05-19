package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second) // Let goroutine finish
	fmt.Println("Main finished")
}

//🧠 What’s a Goroutine?
//A goroutine is a function running concurrently (not necessarily in parallel) with other functions. It's super lightweight (a few KB), thanks to Go’s scheduler magic. You can spawn thousands of them without setting your CPU on fire. 🔥
//
//🚀 How to Start a Goroutine
//
//
//
//go someFunction()
//Just slap go in front of any function call and boom 💥— it’s a goroutine now.

// But remember, it runs in the background, so you might need to wait for it to finish using `sync.WaitGroup` or `time.Sleep()` (not recommended for production).
//🧵 How Are Goroutines Different from Threads?
//Goroutines are like threads but way more efficient. They’re managed by the Go runtime, not the OS. This means you can have thousands of them without breaking a sweat. Threads? Not so much. They’re heavier and managed by the OS, which can lead to resource exhaustion.
// So, if you’re looking to scale your app without turning your server into a sauna, goroutines are your best friend.
// 🐒

//| Feature       | Goroutine      | Thread         |
//| ------------- | -------------- | -------------- |
//| Size          | \~2KB stack    | \~1MB stack    |
//| Creation cost | Extremely low  | Medium–High    |
//| Managed by    | Go runtime     | OS kernel      |
//| Scaling       | Thousands easy | Dozens at most |

//| Communication  | Channels       | Shared memory  |
//| Scheduling     | Go scheduler   | OS scheduler   |
//| Blocking       | Lightweight    | Heavyweight    |
//| Context switch | Fast           | Slow           |
//| Debugging      | Easier         | Harder         |
//| Error handling | Panic/recover   | Try/catch      |
//| Debugging      | Easier         | Harder         |

// Goroutines are scheduled by Go’s green thread scheduler, not by the OS. This makes them super-efficient.
// The Go runtime multiplexes goroutines onto OS threads, so you can have thousands of them without breaking a sweat. Threads? Not so much. They’re heavier and managed by the OS, which can lead to resource exhaustion.
// So, if you’re looking to scale your app without turning your server into a sauna, goroutines are your best friend. 🐒

// ch := make(chan string)
//
// go func() {
//     ch <- "Hello from channel!"
// }()
//
// msg := <-ch
// fmt.Println(msg)

// Goroutines don't live in silos — they talk to each other via channels.

//🧼 Best Practices
//Always monitor goroutines with sync.WaitGroup, not just time.Sleep()

//Use channels to synchronize or share data

//Clean up with context.Context to cancel goroutines

//Avoid capturing loop variables directly (classic gotcha)

//var wg sync.WaitGroup

//wg.Add(1)
//go func() {
//     defer wg.Done()
//     fmt.Println("Inside goroutine")
// }()
//
// wg.Wait()
// fmt.Println("All done")
//Use `sync.Mutex` or `sync.RWMutex` for shared data access
//Mutexes are like locks for your data
//They prevent
