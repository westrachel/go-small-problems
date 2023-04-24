package main

import (
	"fmt"
)

// Write a function that accepts an integer & a string s as parameters
// & returns a string s repeated exactly the specified integer number
// of times.

func main() {
	// Test Cases:
	fmt.Println(repeat(6, "I"))
	fmt.Println(repeat(5, "hello"))
}

func repeat(numTimes int, str string) string {
	var repeated string
	for i := 0; i < numTimes; i++ {
		repeated += str
	}
	return repeated
}
