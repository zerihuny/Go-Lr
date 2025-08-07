package main

import (
	"fmt"
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
}
