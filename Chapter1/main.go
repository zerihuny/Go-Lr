package main

import (
	"fmt"
)

var num1	 = 42    // type inference
var raining bool // explicit type declaration
var num2 int

func main() {
	fmt.Println("Hello, world!")
	displayTime()
	fmt.Println(num1 + num2)
	fmt.Println(raining)
}
