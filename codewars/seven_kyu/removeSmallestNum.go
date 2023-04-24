package main

import (
	"fmt"
)

// Write a function that accepts a list of integers & returns the
// a new list with the smallest value removed
// Reqs:
// > If multiple smallest values, remove the one w/ the smallest index
// > If given empty array/list, return an empty array/list
// > don't change the order of elements & don't mutate the original list

// Algorithm
// i. check if the slice that's passed in has a length of 0; if it does,
//    then return the slice
// ii. declare a smallestNum variable and index that are set to the element
//   at the zeroth index of the argument slice that's passed to the function
// iii. iterate through a slice of given integers and on each iteration:
//     > check if the current element is smaller than the element that
//         the smallestNum variable is set to
//     > if it is smaller, then reassign the smallestNum variable to the current
//        element being iterated over and reassign the variable that's keeping
//        track of the index of the smallest number to the current index
// iv. return slices of the given slice argument depending on what the index of
//    the smallestNum is:
//    > if the index of the smallestNum is 0, then return one slice from
//       the 1st element to the end of the array
//    > if the index of the smallestNum is the last index of the slice, then
//       return one slice from the beginning of the slice argument to the 2nd
//       to last element
//    > if the index of the smallestNum is neither the 0th/last index of the slice,
//      then return 2 slices:
//        > one slice is from the 0th-indexed element up to but excluding the
//           smallestNum element
//        > the other slice is from 1 index after the index of the smallestNum's index
//            to the end of the slice (ie slice[smallestNumIdx+1:]

func main() {
	// Test Cases:
	s1, s2, s3 := []int{77, 42, 54, 33}, []int{-5, 0, -3, 7, -25}, []int{11, 4, 2, 16, 1, 10, 1}
	var s4 []int
	fmt.Println(removeSmallestNum(s1))           // [77, 42, 54]
	fmt.Println(removeSmallestNum(s4))           // []
	fmt.Println(removeSmallestNum(s2))           // [-5, 0, -3, 7]
	fmt.Println(removeSmallestNum(s3))           // [11, 4, 2, 16, 10, 1]
	fmt.Printf("%v, %v, %v, %v", s1, s2, s3, s4) // [77, 42, 54, 33], [-5, 0, -3, 7, -25], [11, 4, 2, 16, 1, 10, 1], []
}

func removeSmallestNum(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	idxOfSmallest := smallestNumAndIdx(slice)
	return sliceExcludingSmallest(slice, idxOfSmallest)
}

func smallestNumAndIdx(slice []int) int {
	smallestNum, idxOfSmallest := slice[0], 0
	for i, num := range slice {
		if num < smallestNum {
			smallestNum = num
			idxOfSmallest = i
		}
	}
	return idxOfSmallest
}

func sliceExcludingSmallest(slice []int, idx int) []int {
	lastIdx := len(slice) - 1
	if idx == 0 {
		return slice[1:]
	} else if idx == lastIdx {
		return slice[:lastIdx]
	} else {
		return mergeSlices(slice[:idx], slice[idx+1:])
	}
}

func mergeSlices(slice1, slice2 []int) []int {
	if len(slice2) == 0 {
		return slice1
	} else {
		var emptySlice []int
		newSlice := addToSlice(emptySlice, slice1)
		return addToSlice(newSlice, slice2)
	}
}

func addToSlice(newSlice, elementsToAdd []int) []int {
	for _, num := range elementsToAdd {
		newSlice = append(newSlice, num)
	}
	return newSlice
}
