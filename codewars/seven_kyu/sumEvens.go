package main

import (
	"fmt"
)

// Write a function accepts a list of integers and returns the sum
// of the all even integers in the list

// Algorithm:
// i. declare and initialize a sum variable that's assigned to 0
// ii. iterate through all the integers in the list and on each iteration
//    > check if the integer is divisible by 2
//    > if it is divisible by 2 then add it to the running total
// iii. return the total

func main() {
	// Test cases:
	fmt.Println(sumEvents([]int{4, 3, 1, 2, 5, 10, 6, 7, 9, 8})) // 30
	var emptySlice []int
	fmt.Println(sumEvents(emptySlice))               // 0
	fmt.Println(sumEvents([]int{-2, 3, 1, 2, 9, 8})) // 8
}

func sumEvents(slice []int) int {
	sum := 0
	for _, num := range slice {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
