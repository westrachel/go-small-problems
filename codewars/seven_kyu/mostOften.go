package main

import "fmt"

// Write a function that accepts a list of numbers and returns
// a description of the count of the most frequent element(s) in
// the list and the corresponding value(s)
// Reqs:
// > input is always an array of integers
// > return string description withn empty slice & count of 0
//    if the given list is empty
// > if there are

// Algorithm:
// i. declare a map whose keys will represent integers in the
//   given list and whose values will represent the count of
//   the number of times the integer key appears in the list
// ii. iterate over the integers in the list and check if the
//   integer has already been added as a key to the map
//  > if it has been added as a key, then add 1 to the value
//     associated with the key
// > if it's not an existing key, then add it as a key w/
//   a corresponding value of 1
// iii. declare mostOften slice & count integer variables to
//  keep track of the element(s) that occur(s) most often in
//  the list and their corresponding counts
// iv. iterate over the keys and values in the map and check if
//  the key's count (its value) is greater than the current number
//  the count variable is assigned to
//    > if the count is greater than the current count then
//       reassign the variable keeping track of the most frequent
//       elements to a slice that contains just the current key
//    > if the count of the current key is equivalent to the most
//      frequent count then add the current element to the slice
//      keeping track of the most frequent elements
// v. return string that interpolates slice of most frequent integers
//   and their count

func main() {
	// Test Cases:
	fmt.Println(mostOften([]int{3, -1, -1, -1, 2, 3, -1, 3, -1, 2, 4, 9, 3})) // The most frequent element(s): [-1] occur 5 times.
	fmt.Println(mostOften([]int{3, 2, 3, 9, -1, 3, 4, -1, 2, 4, 9, 9}))       // The most frequent element(s): [3 9] occur 3 times.
	var emptySlice []int
	fmt.Println(mostOften(emptySlice)) // The most frequent element: [] occurs 0 times.
}

func mostOften(slice []int) string {
	if len(slice) == 0 {
		return "The most frequent element: [] occurs 0 times."
	}
	numCounts := counts(slice)
	var mostOften []int
	var mostOftenCount int
	for num, count := range numCounts {
		if count > mostOftenCount {
			mostOften = []int{num}
			mostOftenCount = count
		} else if count == mostOftenCount {
			mostOften = append(mostOften, num)
		}
	}

	return fmt.Sprintf("The most frequent element(s): %v occur %v times.", mostOften, mostOftenCount)
}

func counts(slice []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range slice {
		_, keyExists := counts[num]
		if keyExists {
			counts[num] += 1
		} else {
			counts[num] = 1
		}
	}
	return counts
}
