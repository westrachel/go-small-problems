package main

import (
	"fmt"
	"math/big"
)

/*
  Write a function that is given a list of integers and finds the
  least common multiple of its arguments.

  Reqs:
   > if no arguments are given, return 1
   > if any argument is 0, return 0
   > all integer inputs are >= 0

 Outer Algorithm:
  i. check if length of the given list is 0; if it is return 1
  ii. iterate through the given list of integers and find the smallest
      number
      > assume the first given integer in the list is the smallest
      > iterate from index 1 to the end of the list and check if the
         current integer being iterated over is smaller than the given integer
      > if it is, then reassign the smallest value to the current element
  iii. declare a current multiple variables that's initialized to the smallest number
  iv. if the smallest number is 0, then return 0
  v. start a loop that iterates from the current multiple and on each iteration:
      > reassign the current multiple to the result of adding the smallest value to the
         value of the current Multiple
      > check if the result is a multiple of all the other given integers
         > iterate through the integers excluding the integer of the smallest num
           > check if the multiple divided by the current integer has a remainder
             of zero
      > if it is then break iteration and return true
      > if it's not then keep iterating
   vi. return the smallest number
*/

func main() {
	// Test cases:
	fmt.Println(LCM(18, 30))        // 90
	fmt.Println(LCM())              // 1
	fmt.Println(LCM(18, 30, 0))     // 0
	fmt.Println(LCM(4, 8, 6))       // 24
	fmt.Println(LCM(18, 12))        // 36
	fmt.Println(LCM(120, 77, 9240)) // 9240
}

func LCM(nums ...int64) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(1)
	}
	smallest, idx := smallestNumAndIdx(nums)
	currMultiple := smallest
	if currMultiple == 0 {
		return big.NewInt(0)
	}

	foundLCM := false
	for !foundLCM {
		currMultiple += smallest
		foundLCM = foundCommonMultiple(currMultiple, idx, nums)
	}

	return big.NewInt(currMultiple)
}

func foundCommonMultiple(multiple, index int64, nums []int64) bool {
	for i := int64(0); i < int64(len(nums)); i++ {
		if i != index {
			if multiple%nums[i] != 0 {
				return false
			}
		}
	}
	return true
}

func smallestNumAndIdx(nums []int64) (int64, int64) {
	smallest, idx := nums[0], 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < smallest {
			smallest = num
			idx = i
		}
	}

	return smallest, int64(idx)
}
