package main

import "fmt"

/*
  Write a function that takes a slice of integers that are comprised
  of all even (odd) integers except for one odd (even) integer. Return
  the integer that is the outlier.

  Data:
  > input: slice of integers that is of length >= 3
  > output: integer outlier

  Algorithm:
  i. subset the first 3 integer elements of the given slice
  ii. declare a count of odds and a count of evens
  iii. iterate over the subset and on each iteration:
       > check if the given integer is divisible by 2 w/ remainder of 0
          > if it is, then add 1 to the count of evens
          > if it's not, then add 1 to the count of odds
  iv. if the count of evens exceeds the count of odds then know that
      the outlier is odd; if the count of odds exceeds the count of
      evens then the outlier is even
  v. iterate over the given slice of integers and on each iteration:
     > if expect outlier to be even, check if the given integer is
        even and if it is break iteration and return the given integer
     > if don't expect outlier to be even, check if the given integer is
       odd and if it is then break iteration and return the given integer
*/

func main() {
	// Test cases:
	fmt.Println(findOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))   // 11
	fmt.Println(findOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))   // 160
	fmt.Println(findOutlier([]int{5, 3, -1, 131, 29, 20, 101, 999})) // 20
	fmt.Println(findOutlier([]int{50, 30, 130, 24, 27, 110}))        // 27
}

func findOutlier(slice []int) int {
	evenFlag := expectEven(slice[:3])

	for _, num := range slice {
		if evenFlag {
			if even(num) {
				return num
			}
		} else {
			if !even(num) {
				return num
			}
		}
	}
	return 0
}

func expectEven(slice []int) bool {
	countOdds, countEvens := 0, 0
	for _, num := range slice {
		if even(num) {
			countEvens += 1
		} else {
			countOdds += 1
		}
	}
	return countEvens < countOdds
}

func even(num int) bool {
	return num%2 == 0
}
