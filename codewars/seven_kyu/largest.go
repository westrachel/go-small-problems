package main

import (
	"fmt"
	"sort"
)

// Write a function that returns a list of the largest n elements
// in a given list of integers

// Algorithm:
// i. check if list is empty or if n is invalid (n <= 0)
//    > if either is true, then return empty list
// ii. if n is greater than or equal to the length of the
//   given list then return the list unsorted
// iii. sort the integers in the given list
//   > sort puts integers in ascending order
// iv. determine the starting index of the slice
//   > if n = 3 and the array has N = 5 elements
//      then the starting index is 2 (= N - n)
//   ex: sorted list is [1, 2, 3, 4, 5]
//     slice[startingIdx:] where startingIdx == 2
//      returns [3, 4, 5]

func main() {
	// Test Cases:
	slice := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(largest(2, slice))                                 // [6, 7]
	fmt.Println(largest(4, []int{59, 106, 23, 0, -5, 111, 34, 7})) // [34,59,106,111]
	fmt.Println(largest(0, []int{7, 4, 23, 19, 0}))                // []
	fmt.Println(largest(-9, []int{7, 4, 23, 19, 0}))               // []
	fmt.Println(largest(8, []int{59, 106, 23, 0, -5, 111, 34, 7})) // [59, 106, 23, 0, -5, 111, 34, 7]

}

func largest(n int, list []int) []int {
	if len(list) == 0 || n <= 0 {
		var emptyList []int
		return emptyList
	} else if n >= len(list) {
		return list
	}

	startingIdx := len(list) - n
	sort.Ints(list)
	return list[startingIdx:]
}
