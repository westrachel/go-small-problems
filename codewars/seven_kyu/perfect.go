package main

import (
	"fmt"
)

// Write a function that accepts an integer and returns a boolean that
// indicates if the number is perfect or not
// Reqs:
// > perfect number = number where divisors (excluding itself) sum to the number itself
// > can assume the input will be >= 0

// Algorithm:
// i. check if number is a small positive, b/c in that case want to avoid iteration,
//   and return false automatically (where n <= 1)
// i. create a list of all the factors of the given integer excluding the given number
//   > initialize a slice that contains 1, b/c 1 is always a factor
//   > iterate from 2 to half of the given number & on each iteration:
//        > check if the current number being iterated over multiplies evenly into
//           the given number
//        > if determine the number is a factor of the given number then add it to
//           the slice of factors
// ii. sum the list of factors
// iii. check if the sum of factors is equal to the given number
//     > if it is return true, otherwise return false

func main() {
	// Test cases:
	fmt.Println(perfect(28)) // true b/c 1 + 2 + 4 + 7 + 14 = 28
	fmt.Println(perfect(25)) // false 1 + 5 = 6
	fmt.Println(perfect(8))  // false 1 + 2 + 4 = 7
	fmt.Println(perfect(1))  // false
	fmt.Println(perfect(0))  // false
}

func perfect(n int) bool {
	if n <= 1 {
		return false
	}
	factors := []int{1}
	for factor := 2; factor <= n/2; factor++ {
		if n%factor == 0 {
			factors = append(factors, factor)
		}
	}
	sum := 0
	for _, factor := range factors {
		sum += factor
	}
	return sum == n
}
