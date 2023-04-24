package main

import "fmt"

/* Write a function that takes a string of parentheses and
returns a boolean indicating if the string contains valid
pairs of parentheses or not.

Reqs:
 > a valid pair must start with an open parentheses "("
 > can start another pair before closing a previous pair
 > empty string should return true even though it doesn't contain parentheses
 > all pairs must be closed at the end of the string
 > strings only contain parentheses

 Algorithm:
 i. declare countOpen and countClosed variables that are initialized to 0
 ii. iterate over the given string and on each iteration, check if
     if the character is an open parentheses or a closed parentheses
      > add 1 to the corresponding count
 iii. check if the value assigned to countClosed is greater than countOpen
     > if it is, stop iterating and return false
 iv. if reach the end, check if countOpen is equal to the countClosed
     > if it's not, then return false
     > if it is return true
*/

func main() {
	// Test cases:
	fmt.Println(ValidParentheses("()"))             // true
	fmt.Println(ValidParentheses(")(()))"))         // false
	fmt.Println(ValidParentheses("("))              // false
	fmt.Println(ValidParentheses("(())((()())())")) // true
	fmt.Println(ValidParentheses(""))               // true
}

func ValidParentheses(parens string) bool {
	countOpen, countClosed := 0, 0
	for _, char := range parens {
		paren := string(char)
		if paren == ")" {
			countClosed++
		} else {
			countOpen++
		}

		if countClosed > countOpen {
			return false
		}
	}

	return countOpen == countClosed
}
