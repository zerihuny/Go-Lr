package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // skip 3
		}
		if i == 5 {
			break // stop at 5
		}
		fmt.Println(i)
	}
}
