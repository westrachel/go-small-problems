package main

import (
	"fmt"
)

// Write a function that accepts an array of integers & returns the
// average rounded down to the nearest integer
// Reqs:
// > array will never be empty or contain non-integer values

func main() {
	// Test Cases:
	arr1, arr2, arr3 := [3]int{77, 88, 90}, [3]int{95, 93, 87}, [5]int{84, 89, 93, 90, 91}
	fmt.Println(average(arr1[:])) // 85
	fmt.Println(average(arr2[:])) // 91
	fmt.Println(average(arr3[:])) // 89
}

func average(slice []int) int {
	denominator, total := len(slice), 0
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}
	return total / denominator
}