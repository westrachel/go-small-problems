package main

import (
	"fmt"
	"strings"
)

/*
Write a function that takes a string and returns a new string
where each character in the new string is "(" if that character
occurs 1x in the given string, or ")" if that character occurs >1
time in the given string.

Reqs:
> duplicates are case insensitive ("I" and "i" rep 2 instances of "i")

Algorithm:
i. declare a map of string keys and integer values

	> keys will keep track of lowercased version of a character and
	   value will represent the first index of that character in the given
	   string

ii. declare a newString variable to push parentheses to
iii. iterate over each character in the given string and on each iteration:

	> check if the lowercased version of the character is already a key in the
	   map
	    > if it is then change the character at that key's index value position
	      in the newString to ")" and add to the end of newString ")"
	    > if the lowercased version of the character is not already a key:
	       > add it as a key and set its index to the current index being
	         iterated over
	       > add to the end of newString "("
iv. return the newString
*/
func main() {
	// Test cases:
	fmt.Println(duplicateEncode("din"))     // "((("
	fmt.Println(duplicateEncode("recede"))  //  "()()()"
	fmt.Println(duplicateEncode("Success")) // ")())())"
	fmt.Println(duplicateEncode("(( @"))    // "))(("

}

func duplicateEncode(word string) string {
	existingChars := make(map[string]int)
	var newString string

	for idx, char := range word {
		lowerChar := strings.ToLower(string(char))
		dupIdx, existsFlag := existingChars[lowerChar]
		if existsFlag {
			newString = reassignIdx(newString, ")", dupIdx)
			newString += ")"
		} else {
			existingChars[lowerChar] = idx
			newString += "("
		}

	}

	return newString

}

func reassignIdx(str, newValue string, index int) string {
	return str[:index] + newValue + str[index+1:]
}
