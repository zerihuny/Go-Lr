package main

import "fmt"

var (
	name string = "John Doe" // string variable
	age  int    = 30         // integer variable
	city string = "New York" // string variable
)

const a string = "Hello, Go!" // constant string
const b int = 100             // constant integer
func usedVariable() {
	// Example of using variables
	fmt.Println("Hello, world!")

	// Displaying the sum of numbers
	fmt.Println(num1 + num2 + int(num3) + int(num5) + int(num6))

	// Displaying the boolean value
	fmt.Println(raining)
	fmt.Println("Name:", name)
	fmt.Println("City:", city)
	// Displaying the complex number
	fmt.Println(num4)
}
