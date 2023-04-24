package main

import (
	"fmt"
	"strings"
)

// Write a function that accepts a string and returns the numerical
// count of vowels in the string
// Reqs:
// vowels = a, e, i, o, u

func main() {
	// Test cases:
	fmt.Println(numVowels("missouri") == 4)
	fmt.Println(numVowels("lemonade") == 4)
	fmt.Println(numVowels("family") == 2)
}

func numVowels(str string) int {
	count := 0
	const Vowels = "AEOIU"
	for i, _ := range str {
		letter := strings.ToUpper(string(str[i]))
		if strings.Contains(Vowels, letter) {
			count += 1
		}
	}

	return count
}
