package main

import "fmt"

func updateName(name *string) {
	*name = "new name"

}

func main() {
	var name = "jam"
	updateName(&name)
	fmt.Println(name) // Output: new nam
}
