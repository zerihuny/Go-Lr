// Package main demonstrates basic variable declarations, type inference, and usage in Go.
// It includes examples of explicit and inferred types, complex numbers, and float types.
// Note: Short variable declarations (:=) and variable concatenation are not allowed at the package level in Go;
// they must be used inside functions or methods.
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

func main() {
	fmt.Println("Hello, world!")
	displayTime()
	fmt.Println(num1 + num2 + int(num3) + int(num5) + int(num6))
	fmt.Println(raining)
	fmt.Println(num4)
	start := time.Now()

	fmt.Printf("%T\n", start)          // Display the type of start variable
	fmt.Println(reflect.TypeOf(start)) // Display the type of start variable using reflect package
	fmt.Println(reflect.ValueOf(start).Kind())
	usedVariable()
}

// displayTime displays the current time in a formatted string.
func displayTime() {
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Current time in UTC:", currentTime.UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("Current time in local timezone:", currentTime.Local().Format("2006-01-02 15:04:05"))
}
