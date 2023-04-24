package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Write a function that accepts a string and returns a boolean
// the indicates if the string is an isogram or not

// Reqs:
// > isogram: contains only unique letters
// > an empty string is an isogram
// > ignore case (so moOse and moose are both not isograms)

// Algorithm:
// i. check if the given string contains any character that is not an
//  alphabetical letter, in which case return false
//    > can use negation regex pattern: [^a-zA-Z] to test for presence
//       of non-alphabetical letter
// ii. iterate through each character and on each iteration:
//    > check if the lowercased version of the letter is already key in
//       a map that keeps track of the counts of each character in the
//       given string
//        > if it is already a key then there's a presence of a duplicate
//          and so we can stop iterating and return false
//        > if it is not a key then initialize it as a key with a count
//          of one
// iii. return true b/c if haven't stopped execution at a prior step to
//   return false, then the string consists of only unique characters

func main() {
	// Test cases:
	fmt.Println(isIsogram("Dermatoglyphics")) // true
	fmt.Println(isIsogram("moose"))           // false
	fmt.Println(isIsogram("aba"))             // false
	fmt.Println(isIsogram("b5an6d"))          // false
	fmt.Println(isIsogram("band"))            // true
	fmt.Println(isIsogram("sunRaY"))          // true
	fmt.Println(isIsogram("sunRaY!"))         // false
}

func isIsogram(str string) bool {
	nonLetter, _ := regexp.MatchString("[^a-zA-Z]", str)
	if nonLetter {
		return false
	}
	counts := make(map[string]int)
	for _, char := range str {
		key := strings.ToUpper(string(char))
		_, exists := counts[key]
		if exists {
			return false
		} else {
			counts[key] = 1
		}
	}

	return true
}
