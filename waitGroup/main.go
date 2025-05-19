package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // We’re launching a new goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Block until all workers call wg.Done()
	fmt.Println("All workers finished ✅")
}

//🧠 What is a WaitGroup?
//Think of a WaitGroup as a bouncer at the door of your main() function.
//
//It:
//
//Waits for a set number of goroutines to finish
//
//Makes sure your program doesn’t exit early
//
//Keeps goroutines accountable like a strict project manager
//
//Without it? Goroutines might be abandoned mid-task like half-eaten pizza 🍕
//
//🔧 Import It First
//go
//Copy
//Edit
//import "sync"
//🛠️ How It Works – Step by Step
//📦 1. Create a WaitGroup
//go
//Copy
//Edit
//var wg sync.WaitGroup
//➕ 2. Tell it how many goroutines you're waiting for
//go
//Copy
//Edit
//wg.Add(1) // Expecting one goroutine
//🔁 3. Inside each goroutine, call wg.Done() when done
//go
//Copy
//Edit
//go func() {
//    defer wg.Done()
//    // ...do work
//}()
//🧍 4. In main(), call wg.Wait() to block until all .Done()s are called
//go
//Copy
//Edit
//wg.Wait()
//✅ Full Example
//go
//Copy
//Edit
//package main
//
//import (
//    "fmt"
//    "sync"
//    "time"
//)
//
//func worker(id int, wg *sync.WaitGroup) {
//    defer wg.Done() // Mark this goroutine as done when it finishes
//    fmt.Printf("Worker %d starting\n", id)
//    time.Sleep(time.Second) // Simulate work
//    fmt.Printf("Worker %d done\n", id)
//}
//
//func main() {
//    var wg sync.WaitGroup
//
//    for i := 1; i <= 3; i++ {
//        wg.Add(1) // We’re launching a new goroutine
//        go worker(i, &wg)
//    }
//
//    wg.Wait() // Block until all workers call wg.Done()
//    fmt.Println("All workers finished ✅")
//}
//🧾 Output:
//bash
//Copy
//Edit
//Worker 1 starting
//Worker 2 starting
//Worker 3 starting
//Worker 2 done
//Worker 1 done
//Worker 3 done
//All workers finished ✅
//(Goroutines may finish in any order — concurrency, baby!)
//
//⚠️ Gotchas
//Gotcha	What happens
//Forget wg.Add(1)	Wait() returns too early
//Forget wg.Done()	Wait() waits forever (deadlock)
//Don’t pass pointer	Changes don’t reflect outside function
//
//✅ Always pass the WaitGroup as a pointer: *sync.WaitGroup
//
//
