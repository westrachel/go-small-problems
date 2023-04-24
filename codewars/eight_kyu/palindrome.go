package main

import (
	"fmt"
	"strings"
)

// Write a function that checks if a given string (case insensitive)
// is a palindrome.

func main() {
	// Test Cases:
	fmt.Println(palindrome("raCecar")) // true
	fmt.Println(palindrome("racecar")) // true
	fmt.Println(palindrome("madame"))  // false
	fmt.Println(palindrome("MaDaM"))   // true
}

func palindrome(str string) bool {
	var flag = true
	for idx1, idx2 := 0, len(str); idx1 < len(str)/2; idx1, idx2 = idx1+1, idx2-1 {
		char1, char2 := str[idx1:idx1+1], str[idx2-1:idx2]
		if strings.ToUpper(char1) != strings.ToUpper(char2) {
			flag = false
		}
	}
	return flag
}
