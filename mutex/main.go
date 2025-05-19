package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

//
//Ahhh, **mutexes** in Go — the cranky gatekeepers that make sure only one goroutine’s allowed in the club at a time. 🕺🔒
//No ticket? No entry.
//Multiple goroutines? Wait in line, fam.
//
//---
//
//## 🧠 What is a Mutex?
//
//A **mutex** (short for *mutual exclusion*) is used to **protect shared resources** from being accessed by **multiple goroutines at the same time**. It ensures that **only one goroutine** can access the **critical section** of code at a time.
//
//> Think of it like a bathroom key in a shared office: one person at a time, or we’re in for a mess 💩
//
//---
//
//## 🔧 Importing Mutex
//
//```go
//import "sync"
//```
//
//---
//
//## 💡 Basic Mutex Example
//
//```go
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//var counter int
//var mu sync.Mutex
//
//func increment(wg *sync.WaitGroup) {
//	mu.Lock()
//	counter++
//	mu.Unlock()
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go increment(&wg)
//	}
//
//	wg.Wait()
//	fmt.Println("Final Counter:", counter)
//}
//```
//
//🧠 **What's Happening:**
//
//* Multiple goroutines call `increment()`.
//* `mu.Lock()` ensures that only **one** goroutine modifies `counter` at a time.
//* `mu.Unlock()` lets the next one in.
//* `WaitGroup` just waits for all to finish.
//
//⛔ Without the mutex? Say hello to **race conditions** and buggy behavior. Your counter might be `3` instead of `5`. 💥
//
//---
//
//## 😎 Go’s Philosophy: Don't use mutexes unless you must
//
//Go really pushes you toward channels for concurrency. The saying is:
//
//> **“Do not communicate by sharing memory; instead, share memory by communicating.”**
//
//But... sometimes you gotta go full mutex because:
//
//* You need **fine-grained control**
//* You're updating **shared state**
//* You’re using **non-channel-based** APIs (e.g., a third-party cache)
//
//---
//
//## 🧃 Read/Write Mutex (RWMutex)
//
//If you're doing **lots of reads but occasional writes**, use `sync.RWMutex`:
//
//```go
//var rw sync.RWMutex
//
//// Read
//rw.RLock()
//// ... read something ...
//rw.RUnlock()
//
//// Write
//rw.Lock()
//// ... write something ...
//rw.Unlock()
//```
//
//This lets **multiple readers** read concurrently — but **writers** get exclusive access.
//
//---
//
//## 🔒 Mutex Do's & Don'ts
//
//✅ Do:
//
//* Always **unlock** after **lock** (or use `defer`).
//* Use `RWMutex` if reads dominate.
//* Keep critical section **short** and **fast**.
//
//🚫 Don't:
//
//* Forget to unlock → **deadlocks**.
//* Lock in multiple places inconsistently → **deadlocks**.
//* Try to unlock without locking → **panic**.
//
//---
//
//## 🔐 Pro Tip: `defer` is your best friend
//
//```go
//mu.Lock()
//defer mu.Unlock()
//// safe and sound operations here
//```
//
//No matter how the function exits, the mutex will unlock.
//
//---
//
//## 🔨 Commit Message Ideas
//
//```bash
//fix: added mutex to prevent concurrent counter updates
//```
//
//or poetic mode:
//
//```bash
//feat(sync): added mutex, because chaos is a ladder and I need a lock
//```
//
//---
//
//Want a mutex vs channel showdown? Or build a cache with mutexes?
//Let’s rumble 🥊— just say the word.
//
