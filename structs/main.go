//🧱 What is a struct in Go?
//A struct (short for “structure”) is a composite data type that groups together variables under one name. These variables are called fields.
//
//Think of it like a class in OOP — but without methods tied in (methods get attached separately).

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Uday", Age: 30}
	fmt.Println(p.Name) // Uday
	fmt.Println(p.Age)  // 30
}

// Structs are useful for:
// 1. Grouping related data together.
// 2. Creating complex data types.
// 3. Defining custom data structures.
// 4. Implementing data models.
// 5. Creating APIs.
// 6. Representing real-world entities.
// 7. Storing configuration data.
// 8. Creating data transfer objects (DTOs).

//func birthday(p *Person) {
//    p.Age++
//}
//
//func main() {
//    p := Person{Name: "Uday", Age: 30}
//    birthday(&p)
//    fmt.Println(p.Age) // 31
//}
//// Structs are value types, so when you pass them to functions, they are copied. To modify the original struct, you need to pass a pointer to it.
//// This is done using the & operator to get the address of the struct and the * operator to dereference it.
////
//// Structs can also have methods, which are functions that operate on the struct.
//// This is done by defining a receiver type (the struct type) before the function name. The receiver type can be a value or a pointer.
//// This allows you to define methods that can modify the struct or access its fields.
////
//// Methods are defined like this:
//// func (r ReceiverType) MethodName(parameters) returnType {
////     // method body
//// }
////
//// Example of a method on a struct:
//package main
//import "fmt"
//type Person struct {
//	Name string
//	Age  int
//}
//func (p *Person) Birthday() {
//	p.Age++
//}
//func main() {
//	p := Person{Name: "Uday", Age: 30}
//	p.Birthday()
//	fmt.Println(p.Age) // 31
//}
//// In this example, the Birthday method modifies the Age field of the Person struct.
//// The receiver type is *Person, which means the method can modify the original struct.
//// If you want to define a method that does not modify the struct, you can use a value receiver instead:
//package main
//import "fmt"
//type Person struct {
//	Name string
//	Age  int
//}
//func (p Person) Greet() {
//	fmt.Println("Hello, my name is", p.Name)
//}
//func main() {
//	p := Person{Name: "Uday", Age: 30}
//	p.Greet() // Hello, my name is Uday
//}
//// In this example, the Greet method does not modify the struct, so we use a value receiver.
//// This means that the method receives a copy of the struct, and any modifications made to it will not affect the original struct.
//// Structs can also be nested, meaning you can have a struct as a field in another struct. This allows you to create complex data structures.
//// For example:
//package main
//import "fmt"
//type Address struct {
//	City    string
//	State   string
//}
//type Person struct {
//	Name    string
//	Age     int
//	Address Address
//}
//func main() {
//	p := Person{
//		Name: "Uday",
//		Age:  30,
//		Address: Address{
//			City:  "New York",
//			State: "NY",
//		},
//	}
//	fmt.Println(p.Name)    // Uday
//	fmt.Println(p.Age)     // 30
//	fmt.Println(p.Address)  // {New York NY}
//	fmt.Println(p.Address.City) // New York
//}
//// In this example, the Person struct has an Address field, which is another struct.
//// This allows you to group related data together and create complex data structures.
//// Structs can also be used to implement interfaces, which are a way to define a contract for types.
//// An interface defines a set of methods that a type must implement. This allows you to create flexible and reusable code.
//// For example:
//package main
//import "fmt"
//type Shape interface {
//	Area() float64
//}
//type Rectangle struct {
//	Width  float64
//	Height float64
//}
//func (r Rectangle) Area() float64 {
//	return r.Width * r.Height
//}
//type Circle struct {
//	Radius float64
//}
//func (c Circle) Area() float64 {
//	return 3.14 * c.Radius * c.Radius
//}
//func main() {
//	r := Rectangle{Width: 10, Height: 5}
//	c := Circle{Radius: 7}
//	fmt.Println("Rectangle Area:", r.Area()) // Rectangle Area: 50
//	fmt.Println("Circle Area:", c.Area())     // Circle Area: 153.86
//}
//// In this example, we define a Shape interface with an Area method. The Rectangle and Circle structs implement the Area method, allowing us to calculate their areas.
//// This demonstrates how structs can be used to create flexible and reusable code by implementing interfaces.
//// Structs can also be used to create custom data types, which can help improve code readability and maintainability.
//// For example, you can create a custom data type for a specific purpose, such as representing a date or a currency.
//// This can help make your code more self-documenting and easier to understand.
//// For example:
//package main
//import "fmt"
//type Date struct {
//	Day   int
//	Month int
//	Year  int
//}
//func (d Date) String() string {
//	return fmt.Sprintf("%02d/%02d/%04d", d.Day, d.Month, d.Year)
//}
//func main() {
//	d := Date{Day: 15, Month: 8, Year: 2023}
//	fmt.Println(d) // 15/08/2023
//}
//// In this example, we define a Date struct with Day, Month, and Year fields. We also define a String method to format the date as a string.
//// This allows us to create a custom data type for representing dates, making our code more self-documenting and easier to understand.
//// Structs can also be used to create data transfer objects (DTOs), which are used to transfer data between different layers of an application.
//// DTOs are often used in web applications to transfer data between the server and the client.
//// They can help improve code readability and maintainability by providing a clear structure for the data being transferred.
//// For example:
//package main
//import "fmt"
//type UserDTO struct {
//	Username string
//	Email    string
//	Age      int
//}
//
//func main() {
//	u := UserDTO{Username: "Uday", Email: "uday@example.com", Age: 30}
//	fmt.Println(u)
//}// In this example, we define a UserDTO struct with Username, Email, and Age fields. This allows us to create a clear structure for the data being transferred between the server and the client.
//// Structs can also be used to create APIs, which are a way to expose functionality to other applications or services.
//// APIs often use structs to define the data being sent and received, making it easier to understand and work with the data.
//// For example:
//package main
//import "fmt"
//type User struct {
//	Username string
//	Email    string
//	Age      int
//}
//type UserAPI struct {
//	Users []User
//}
//func (api *UserAPI) AddUser(user User) {
//	api.Users = append(api.Users, user)
//}
//func (api *UserAPI) GetUsers() []User {
//	return api.Users
//}
//func main() {
//	api := UserAPI{}
//	api.AddUser(User{Username: "Uday", Email: "uday@example.com", Age: 30})
//	fmt.Println(api.GetUsers())
//}
//// In this example, we define a User struct and a UserAPI struct that contains a slice of User structs. The UserAPI struct has methods to add and get users.
//// This allows us to create an API for managing users, making it easier to work with the data.
//// Structs can also be used to create configuration data, which is often used to store settings or options for an application.
//// Configuration data can be stored in structs, making it easier to manage and work with the data.
//// For example:
//package main
//import "fmt"
//type Config struct {
//	Host     string
//	Port     int
//	Username string
//	Password string
//}
//
//func main() {
//	c := Config{Host: "localhost", Port: 8080, Username: "admin", Password: "password"}
//	fmt.Println(c)
//}
//// In this example, we define a Config struct with Host, Port, Username, and Password fields. This allows us to create a clear structure for the configuration data, making it easier to manage and work with.
//// Structs can also be used to create custom data types for specific purposes, such as representing a currency or a measurement.
//// This can help improve code readability and maintainability by providing a clear structure for the data being represented.
//
//func birthday(p *Person) {
//    p.Age++
//}
//
//func main() {
//    p := Person{Name: "Uday", Age: 30}
//    birthday(&p)
//    fmt.Println(p.Age) // 31
//}
//
//
//func (p Person) Greet() {
//    fmt.Println("Hello, my name is", p.Name)
//}
//
//func (p *Person) HaveBirthday() {
//    p.Age++
//}
//
//
//Structs in Go
//Group multiple fields into one type
//
//Behave like value types (copy on assignment)
//
//Methods can be attached (via value or pointer receiver)
//
//Support embedding instead of inheritance
//
//Can use tags for encoding/decoding
//
//No constructors — factory functions instead
//
//

//Alright, let's break this down like it’s a TikTok tutorial for Go structs — short, snappy, and no BS. 🎬
//
//---
//
//### 🔨 The Struct in Question:
//
//```go
//type User struct {
//    ID   int    `json:"id"`
//    Name string `json:"name"`
//}
//```
//
//This line of Go code is defining a **struct type** named `User`, which is a custom data type. Think of it like a blueprint for an object. Here's what each piece means:
//
//---
//
//### 🧱 `type User struct { ... }`
//
//* You're **defining a new struct type** called `User`.
//* It’s made up of **fields**: `ID` and `Name`.
//* You can now create variables of type `User` to represent a user in your program.
//
//---
//
//### 🔑 The Fields:
//
//#### `ID   int`
//
//* Field name: `ID`
//* Field type: `int` (an integer)
//* It will hold a numeric user identifier (e.g. `1`, `42`, `999`).
//
//#### `Name string`
//
//* Field name: `Name`
//* Field type: `string`
//* Holds the user’s name (e.g. `"Uday"` or `"Neo"`).
//
//---
//
//### 🏷️ What’s up with the Backticks?
//
//```go
//`json:"id"`
//```
//
//These are called **struct tags**, and they’re actually really powerful.
//
//#### 🔍 Why use `json:"id"`?
//
//* Go’s `encoding/json` package uses **reflection** to know how to map JSON keys to struct fields.
//* Without the tag, Go defaults to using the **field name** (case-sensitive).
//* With the tag, you're saying:
//
//  > “When encoding/decoding JSON, map the JSON key `id` to the Go field `ID`.”
//
//#### Example in action:
//
//```go
//user := User{ID: 101, Name: "Uday"}
//jsonData, _ := json.Marshal(user)
//fmt.Println(string(jsonData))
//// Output: {"id":101,"name":"Uday"}
//```
//
//If you **don’t use** the tag, the JSON would look like:
//
//```json
//{"ID":101,"Name":"Uday"}
//```
//
//Which isn't idiomatic in JSON land (JSON uses lowercase snake\_case or camelCase, while Go uses PascalCase for exported fields).
//
//---
//
//### 🧪 TL;DR Breakdown
//
//| Part               | Meaning                                                                                 |
//| ------------------ | --------------------------------------------------------------------------------------- |
//| `type User struct` | Declares a new struct named `User`                                                      |
//| `ID int`           | Field named `ID` of type `int`                                                          |
//| `Name string`      | Field named `Name` of type `string`                                                     |
//| `` `json:"id"` ``  | Tells the JSON encoder/decoder to use `"id"` instead of `"ID"` in the JSON output/input |
//
//---
//
//### 🎁 Bonus: Field Visibility
//
//* `ID` and `Name` are capitalized = **exported**, which means:
//
//  * Accessible **outside** the current package.
//  * Necessary for JSON encoding/decoding to work!
//
//If you wrote:
//
//```go
//type User struct {
//    id   int    `json:"id"`
//    name string `json:"name"`
//}
//```
//
//Then even with the JSON tags, the encoding wouldn’t work properly — Go's JSON encoder only works on **exported** fields (those with capital letters).
