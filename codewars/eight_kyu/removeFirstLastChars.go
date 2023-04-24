package main

import (
	"fmt"
)

// Write a function that removes the first & last characters
// of a given string. Given strings are >= 2 characters

func main() {
	// Test Cases:
	fmt.Println(removeFirstLastChars("he"))     // ""
	fmt.Println(removeFirstLastChars("madam"))  // "ada"
	fmt.Println(removeFirstLastChars("orange")) // "rang"
}

func removeFirstLastChars(str string) string {
	return str[1 : len(str)-1]
}
