package main

func main() {
	// condtional examples
	// if statement
	if true {
		println("This is true")
	} else {
		println("This is false")
	}
	// if with initialization
	if x := 10; x > 5 {
		println("x is greater than 5")
	} else {
		println("x is less than or equal to 5")
	}
	// switch statement
	switch x := 10; {
	case x < 5:
		println("x is less than 5")
	case x == 10:
		println("x is equal to 10")
	default:
		println("x is greater than 10")

	}
	// switch with initialization
	switch y := 20; {
	case y < 10:
		println("y is less than 10")
	case y == 20:
		println("y is equal to 20")
	default:
		println("y is greater than 20")
	}
	// type switch
	var i interface{} = "Hello"
	switch i.(type) {
	case int:
		println("i is an int")
	case string:
		println("i is a string")
	default:
		println("i is of a different type")
	}

}
