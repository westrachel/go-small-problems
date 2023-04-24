package main

import (
	"fmt"
)

// Write a function that accepts a list of an odd-number of integers
// that contains duplicates & 1 unique integer & returns the singular
// unique number
// Reqs:
// > slice input length is always odd and is >= 3
// > there is only one type of duplicate, not multiple sets of duplicates

// Algorithm:
// i. declare num1 and num2 variables and initialize them to the element at
//    the 0th index position of the given slice and 0
// ii. declare count1 and count2 variables that keep track of the counts of
//    occurrences of num1 and num2
// iii. iterate over the remainder of the slice (skip the element at the 0th
//    index, so that we don't automatically set the 0th indexed element as
//    the duplicate) on each iteration
//    > check if the element is equal to the value assigned to num1
//      > if it is, then add 1 to the value that count1 is assigned to
//    > if the element is not equal to the value assigned to num1
//        then reassign num2 to the current element and add1 to
//        value assigned to count2
//    > check if either of the values assigned to the count variables
//       exceed 1, if they do then reassign the unique variable to the
//       opposite number
//         > this means that if count1 is greater than 1 then reassign
//           unique to num2
//         > end iteration early by:
//           > adding break statement after reassinging count2 to 2 b/c
//             then we know num1 must be the unique value
//           > adding break statement after count1 exceeds 1 and count2 is
//             nonzero
//              > count2 must be nonzero to ensure that we've found the
//                  the unique value b/c could have slice that has a lot
//                  of duplicates in a row from the start

func main() {
	// Test cases:
	s1, s2 := []int{1, 1, 2}, []int{17, 17, 3, 17, 17, 17, 17}
	s3, s4 := []int{5, 4, 4, 4, 4, 4, 4}, []int{17, 17, 17, 17, 3, 17, 17}
	fmt.Println(stray(s1)) // 2
	fmt.Println(stray(s2)) // 3
	fmt.Println(stray(s3)) // 5
	fmt.Println(stray(s4)) // 3
}

func stray(slice []int) int {
	num1, num2 := slice[0], 0
	count1, count2 := 1, 0
	var unique int

	for _, num := range slice[1:] {
		if num == num1 {
			count1 += 1
		} else {
			num2 = num
			count2 += 1
		}

		if count1 > 1 {
			unique = num2
			if count2 > 0 {
				break
			}
		} else if count2 > 1 {
			unique = num1
			break
		}
	}
	return unique
}
