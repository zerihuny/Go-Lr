package main

import (
	"fmt"
	"reflect"
	"time"
)

var num1 = 42    // type inference
var raining bool // explicit type declaration
var num2 int
var num3 = 3.14   // type inference
var num4 = 1 + 2i // complex number
var num5 = 1.618  // float64
var num6 float32

// firstName := "John" // short variable declaration
// lastName := "Doe"   // short variable declaration
// Note: The above short variable declarations will not work at package level in Go.
// They should be used inside a function or a method.
// var fullName = firstName + " " + lastName // concatenation
// Note: The above concatenation will not work at package level in Go.
// It should be done inside a function or a method.
// firstName, lastName, age := "John", "Doe", 30 // multiple short variable declarations
// Note: The above multiple short variable declarations will not work at package level in Go.
// They should be used inside a function or a method.

func main() {
	fmt.Println("Hello, world!")
	displayTime()
	fmt.Println(num1 + num2 + int(num3) + int(num5) + int(num6))
	fmt.Println(raining)
	fmt.Println(num4)
	start := time.Now()

	fmt.Printf("%T\n", start)          // Display the type of start variable
	fmt.Println(reflect.TypeOf(start)) // Display the type of start variable using reflect package
	usedVariable()
}
