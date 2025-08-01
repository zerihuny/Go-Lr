package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello this is my first Go program!")
	result := add(5, 3)
	fmt.Println("The sum of 5 and 3 is:", result)
	fmt.Println("Goodbye!")

}
