package main

import (
	"fmt"
	"regexp"
)

// Write a function that takes a string argument and returns a list
// of integers that informs how many uppercase letters, lowercase,
// numbers, and special characters there are in the string

// Algorithm:
// i. declare 4 variables that all start at zero that keep track of
//     the counts of different types of characters; initialize them to 0
// ii. iterate through the string and on each iteration check if the individual
//     string character matches via different regex patterns:
//        an uppercase string: [A-Z]
//        a lowercase string: [a-z]
//        a numerical string digit [0-9]
//        none of the above, in which case it's a special character
// iii. return a slice that contains each of the integer values
//     order to return values in:
//       count uppercase, count lowercase, count integer, count specialChars

func main() {
	// Test cases:
	fmt.Println(specialCharCounts("*'&ABCDabcde12345")) // [4,5,5,3]
}

func specialCharCounts(str string) []int {
	upper, lower, num, other := 0, 0, 0, 0
	for _, char := range str {
		if regexp.MustCompile(`[A-Z]`).FindString(string(char)) != "" {
			upper += 1
		} else if regexp.MustCompile(`[a-z]`).FindString(string(char)) != "" {
			lower += 1
		} else if regexp.MustCompile(`[0-9]`).FindString(string(char)) != "" {
			num += 1
		} else {
			other += 1
		}
	}
	return []int{upper, lower, num, other}
}
