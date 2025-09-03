package main

import "fmt"

func main() {
	result := DecisionExample(85) // Returns "B"
	fmt.Println(result)
	logicals()

}

func DecisionExample(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}
func logicals() {
	num := 10
	condition := num > 3 && num < 29
	fmt.Println(condition)
	// import "fmt"
}
