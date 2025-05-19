package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//Ah, **files in Go** — where your code stops talking to itself and starts talking to the outside world 📁🧃
//Reading from files, writing to files, appending logs, all that grown-up stuff — Go's `os` and `io/ioutil` (now `os` and `io` or `os` and `io/fs` in modern Go) packages got your back.
//
//Let’s break it down like a DJ spinning bytes. 🎧
//
//---
//
//## 🔨 Basic File Operations in Go
//
//### 1. **Read a File (like, just open and read everything)**
//
//```go
//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	data, err := os.ReadFile("example.txt")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(data))
//}
//```
//
//🧠 `os.ReadFile()` returns the **entire content** as `[]byte`.
//
//---
//
//### 2. **Write to a File (overwrites the file)**
//
//```go
//package main
//
//import (
//	"os"
//)
//
//func main() {
//	err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
//	if err != nil {
//		panic(err)
//	}
//}
//```
//
//💡 Permission `0644` means: owner can read/write, others can only read.
//
//---
//
//### 3. **Append to a File**
//
//```go
//package main
//
//import (
//	"os"
//)
//
//func main() {
//	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	if _, err := f.WriteString("Adding another line\n"); err != nil {
//		panic(err)
//	}
//}
//```
//
//This is your logger, your journal, your “dear diary” moment. ✍️
//
//---
//
//### 4. **Read File Line by Line (Scanner)**
//
//```go
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	file, err := os.Open("example.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}
//
//	if err := scanner.Err(); err != nil {
//		panic(err)
//	}
//}
//```
//
//This is **memory efficient** — no need to load the whole file into memory.
//
//---
//
//### 5. **Create a File**
//
//```go
//f, err := os.Create("newfile.txt")
//if err != nil {
//	panic(err)
//}
//defer f.Close()
//
//f.WriteString("Fresh file vibes 🎉")
//```
//
//`os.Create()` overwrites the file if it exists. No second chances 😅
//
//---
//
//## 🧪 File Modes Cheat Sheet
//
//| Flag          | Meaning                         |
//| ------------- | ------------------------------- |
//| `os.O_RDONLY` | Read only                       |
//| `os.O_WRONLY` | Write only                      |
//| `os.O_RDWR`   | Read + Write                    |
//| `os.O_APPEND` | Append to file                  |
//| `os.O_CREATE` | Create if not exists            |
//| `os.O_TRUNC`  | Truncate (clear file) if exists |
//
//You can **bitwise OR (`|`)** them together like `os.O_CREATE|os.O_WRONLY`.
//
//---
//
//## 📎 Commit Message Ideas
//
//```bash
//feat: added file read/write functionality
//```
//
//or nerdy flair:
//
//```bash
//feat(fs): gave Go the power to speak to disk ✍️📂
//```
//
//---
//
//## 🔥 Want More?
//
//* JSON → file?
//* File uploads in HTTP servers?
//* Temporary files and directories?
//* Stream processing big CSV files?
//
//Say the word and I’ll drop more file-flavored Go knowledge faster than `os.Remove("temp.txt")` 🧨
