package main

import (
	"fmt"
	"math"
)

// Write a function that determines if a number is square

func main() {
	// Test Cases:
	fmt.Println(isSquare(-1)) // false
	fmt.Println(isSquare(26)) // false
	fmt.Println(isSquare(4))  // true
	fmt.Println(isSquare(25)) // true
}

func isSquare(num int) bool {
	squareRoot := math.Sqrt(float64(num))
	if squareRoot == math.Round(squareRoot) {
		return true
	} else {
		return false
	}
}
