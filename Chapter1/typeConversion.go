package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	// fmt.Println("You entered:", input)
	age, err := strconv.Atoi(input)
	if err != nil { // an error occurred
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}

func anotherOne() {

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	// fmt.Println("You entered:", input)
	age, err := strconv.Atoi(input)
	if err != nil { // an error occurred
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}

func yetAnotherOne() {

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	// fmt.Println("You entered:", input)
	age, err := strconv.Atoi(input)
	if err != nil { // an error occurred
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}

// another function to demonstrate type conversion
func typeConversionDemo() {

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	// fmt.Println("You entered:", input)
	age, err := strconv.Atoi(input)
	if err != nil { // an error occurred
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}
}
