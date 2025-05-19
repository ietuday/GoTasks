package main

import "fmt"

func sayHi(ch chan string) {
	ch <- "Hello from goroutine!"
}

func main() {
	ch := make(chan string) // make a channel

	go sayHi(ch) // run goroutine

	msg := <-ch // receive from channel (blocks until received)
	fmt.Println(msg)
}

//
// Let’s gooo 🏁 — you're about to master **Channels in Go**.
// Think of channels as the **walkie-talkies of goroutines** — they let your concurrently running Go functions talk to each other safely and efficiently.
//
// ---
//
// ## 🧠 What’s a Channel in Go?
//
// A **channel** is a **typed conduit** through which goroutines can communicate and synchronize by **sending and receiving values**.
//
// ```go
// ch := make(chan int)
// ```
//
// You can then:
//
// * **Send** with: `ch <- value`
// * **Receive** with: `value := <-ch`
//
// Boom. That’s your first telepathic link between two goroutines 🔗
//
// ---
//
// ## 📦 Simple Example
//
// ```go
// package main
//
// import "fmt"
//
// func sayHi(ch chan string) {
// 	ch <- "Hello from goroutine!"
// }
//
// func main() {
// 	ch := make(chan string) // make a channel
//
// 	go sayHi(ch)            // run goroutine
//
// 	msg := <-ch             // receive from channel (blocks until received)
// 	fmt.Println(msg)
// }
// ```
//
// 🧾 **Output:**
//
// ```
// Hello from goroutine!
// ```
//
// The `main` function will **pause at `<-ch>`** until the goroutine sends a message. Clean and race-free!
//
// ---
//
// ## 🛠 Types of Channels
//
// ### 1. **Unbuffered Channel** (default)
//
// * Both send and receive block until the other side is ready.
// * Great for tight coordination.
//
// ```go
// ch := make(chan int)
// ```
//
// ### 2. **Buffered Channel**
//
// * Can hold some number of values before blocking.
//
// ```go
// ch := make(chan int, 3)
//
// ch <- 1
// ch <- 2
// ch <- 3 // still OK
// // ch <- 4 // blocks because buffer is full
// ```
//
// ---
//
// ## 🔄 Range + Close on Channels
//
// You can **close** a channel when you're done sending:
//
// ```go
// close(ch)
// ```
//
// And you can **iterate** over a closed channel:
//
// ```go
// for val := range ch {
// 	fmt.Println(val)
// }
// ```
//
// Only the **sender** should call `close()` — receivers just receive and vibe 🧘
//
// ---
//
// ## ⚡ Select Statement = Channel Multiplexing
//
// You can listen to **multiple channels at once** using `select`:
//
// ```go
// select {
// case msg1 := <-ch1:
//     fmt.Println("Got", msg1)
// case msg2 := <-ch2:
//     fmt.Println("Got", msg2)
// default:
//     fmt.Println("No messages received")
// }
// ```
//
// This is Go’s version of non-blocking channel ops. Smooth AF.
//
// ---
//
// ## 🧪 Real-World Use Case: Worker Pool
//
// ```go
// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, j)
// 		results <- j * 2
// 	}
// }
//
// func main() {
// 	jobs := make(chan int, 5)
// 	results := make(chan int, 5)
//
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
//
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
//
// 	for a := 1; a <= 5; a++ {
// 		fmt.Println(<-results)
// 	}
// }
// ```
//
// This right here? 💥 That’s how you process jobs concurrently and get results without losing your mind.
//
// ---
//
// ## ⚠️ Gotchas
//
// | Mistake                | Consequence                  |
// | ---------------------- | ---------------------------- |
// | Send on closed channel | **Panic** 😱                 |
// | Forget to receive      | **Deadlock** 💀              |
// | Forget to close        | Range loops block forever 🚫 |
//
// ---
//
//
// ---
//
// Want to dive into:
//
// * **Select + timeout/cancellation with `context`**
// * **Pipeline patterns** using channels
// * **Unidirectional channels** for function safety
//
