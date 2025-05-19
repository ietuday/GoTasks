//```🧠 What’s an Enum?```
//```In other languages, an enum is a named list of constants, typically used to represent a fixed set of related values.```
//``````
//```Example in TypeScript:```
//``````
//```enum Color {```
//```  Red,```
//```  Green,```
//```  Blue```
//```}```
//```🛠️ Enums in Go using const + iota```
//``````
//```type Color int```
//``````
//```const (```
//```    Red Color = iota```
//```    Green```
//```    Blue```
//```)```
//```💡 Explanation:```
//```iota starts at 0 and auto-increments with each line.```
//``````
//```So Red = 0, Green = 1, Blue = 2.```
//``````
//```You just created an enum-like set of values.```
//``````
//```📦 Why use a named type (type Color int)?```
//```Adds type safety: now only Color values are allowed where expected.```
//``````
//```Helps with attaching methods to your "enum."```
//``````
//```🗣️ Adding a String() method```
//```To make your enum printable (instead of just showing numbers), add a String() method:```
//``````
//```func (c Color) String() string {```
//```    switch c {```
//```    case Red:```
//```        return "Red"```
//```    case Green:```
//```        return "Green"```
//```    case Blue:```
//```        return "Blue"```
//```    default:```
//```        return "Unknown"```
//```    }```
//```}```
//``````
//```fmt.Println(Red)   // Red```
//```fmt.Println(Color(2)) // Blue```
//``````

package main

import (
	"fmt"
)

type Color int

const (
	Red Color = iota
	Green
	Blue
)

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
func main() {
	fmt.Println(Red)      // Red
	fmt.Println(Color(2)) // Blue
}
