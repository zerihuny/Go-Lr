package main

import (
	"fmt"
	"time"
)

func displayTime() {
	fmt.Println(time.Now())
}

// Note: The displayTime function is commented out in main.go to avoid compilation errors
// when running the program, as it is not defined in main.go.
//func main() {	
//	fmt.Println("Hello, world!")
//	displayTime()
// }
// Uncomment the above main function to see the displayTime function in action.
// The displayTime function prints the current time to the console.
// This function can be used to demonstrate how to work with time in Go.	