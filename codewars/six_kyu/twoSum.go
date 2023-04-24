package main

import "fmt"

/* Write a function that takes an list of integers and a target number. It should
  return a list that contains indices of the 2 integers in the lest that when added
  together sum to the target value.

  Reqs:
  > given lists are >= 2 items in length

Algorithm:
  i. declare an empty map that will store integer keys and values
  ii. iterate over the given list and on each iteration:
      > find the result of subtracting the current integer being
        iterated over from the target number
      > check if the result is already a key in the map
          > if is, then return that key's value (which is its index)
           and the current integer's index in a list
             >> return the index of the key in the map in the list
               before the current integer's index, so that they are
               in ascending order
      > push the current integer to the map as a key and assign
        its index as its value
  iii. if there's always a solution then will return while iterating
     over the given list; if there's no solution then will return
     a list containing 2 zeros

*/

func main() {
	// Test cases:
	fmt.Println(twoSum([]int{1, 2, 3}, 4))      // returns [2]int{0, 2}
	fmt.Println(twoSum([]int{1, 3, 5, 10}, 13)) // returns [2]int{1, 3}
	fmt.Println(twoSum([]int{1, 1}, 2))         // returns [2]int{0, 1}
	fmt.Println(twoSum([]int{-4, 0, 4}, 4))     // returns [2]int{1, 2}
}

func twoSum(nums []int, target int) []int {
	pairs := make(map[int]int)

	for idx, num := range nums {
		difference := target - num
		valueIdx, existsFlag := pairs[difference]

		if existsFlag {
			return []int{valueIdx, idx}
		} else {
			pairs[num] = idx
		}
	}

	return []int{0, 0}
}