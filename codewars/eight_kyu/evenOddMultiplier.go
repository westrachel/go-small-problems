package main

import (
	"fmt"
)

// Write a function that accepts an integer & multiplies it
// by eight if it is an even number and by nine otherwise

func main() {
	// Test Cases:
	fmt.Println(evenOddMultiplier(6)) // 48
	fmt.Println(evenOddMultiplier(5)) // 45
}

func evenOddMultiplier(num int) int {
	if num%2 == 0 {
		return num * 8
	} else {
		return num * 9
	}
}
