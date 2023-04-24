package main

import (
	"fmt"
)

/*
  Write a function that takes a list of integers and returns a list
  of the same size with all zeros moved to the end of the list

  Algorithm:
  i. declare an empty list, newList
  ii. declare a variable countZeros initialized to 0
  iii. iterate over the given list and on each iteration:
     > check if the integer is not zero, if it's not then add it to
      the newList created in (i)
     > if the integer is zero then add 1 to the value assigned to countZeros
  iv. iterate from 1 to the value of countZeros and on each iteration add
    a 0 to newList
  v. return newList
*/

func main() {
	// Test cases:
	fmt.Println(moveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
	fmt.Println(moveZeros([]int{0, 0, 5, 6}))                   // returns []int{5, 6, 0, 0}
}

func moveZeros(list []int) []int {
	var newList []int
	countZeros := 0

	for _, num := range list {
		if num != 0 {
			newList = append(newList, num)
		} else {
			countZeros += 1
		}
	}
	return addZeros(newList, countZeros)
}

func addZeros(allList []int, numZeros int) []int {
	for count := 1; count <= numZeros; count++ {
		allList = append(allList, 0)
	}
	return allList
}
