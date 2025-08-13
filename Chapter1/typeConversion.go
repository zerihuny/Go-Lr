package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Println("Please enter your age: ")
	fmt.Scanf("%s", &input)
	// fmt.Println("You entered:", age)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting age:", err)
	} else {
		fmt.Println("Converted age:", age)
	}
}
