package main

import "fmt"

// Write a function that accepts factor and upperBound integers and
// finds the largest integer N, such that:
//   > factor is a factor of N
//   > N <= upperBound
//   > N > 0

// Reqs:
// > factor & upperBound are both greater than 0

// Algorithm:
// > declare a maxMultiple variable that's equal to the upperBound
// > start a loop that iterates backwards from maxMultiple down to
//    factor and on each iteration:
//     > check if maxMultiple divided by factor has a remainder of 0
//     > if it has a 0 remainder then stop iteration and return the
//         maxMultiple

func main() {
	// Test cases:
	fmt.Println(maxMultiple(2, 7))    // (6)
	fmt.Println(maxMultiple(10, 50))  // (50)
	fmt.Println(maxMultiple(37, 200)) // (185)
}

func maxMultiple(factor, upperBound int) int {
	for num := upperBound; num >= factor; num-- {
		//fmt.Println(num)
		//fmt.Println(factor)
		//fmt.Println(num % factor)

		if num%factor == 0 {
			return num
		}
	}
	return factor
}
