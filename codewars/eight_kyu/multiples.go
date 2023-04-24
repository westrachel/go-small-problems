package main

import "fmt"

// Write a function that takes an integer & a limit & returns a list of
// multiples of the number up to the limit. If the limit is a multiple of
// the integer, it should be included in the return.
// Reqs:
// > only positive integer inputs
// > limit is greater than the integer always

// Test cases:
// (2, 6) -> [2, 4, 6]
// (3, 13) -> [3, 6, 9, 12]
// (5, 37) -> [5, 10, 15, 20, 25, 30, 35]

func main() {
	fmt.Printf("Multiples: %v", multiples(2, 6))
	fmt.Printf("Multiples: %v", multiples(3, 13))
	fmt.Printf("Multiples: %v", multiples(5, 37))
}

func multiples(integer, limit int) []int {
	arr := [1]int{integer}
	slice := arr[:]
	for num := 2 * integer; num <= limit; num += integer {
		slice = append(slice, num)
	}
	return slice
}
