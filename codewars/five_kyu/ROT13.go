package main

import (
	"fmt"
	"strings"
)

/*
  Write a function that unscrambles alphabetical characters in a given
  string that were originally scrambled with ROT13.

  Reqs:
  > ROT13 shifts characters 13 over in the alphabet
    a - b - c - d - e - f - g - h - i - j - k - l - m - n
   [0] 						       [13]
    r - s - t - u - v - w - x - y - z - a - b - c - d - e
   [17]						       [4]
	newIdx := 17 + 13 = 30 - 26 = 5

  Algorithm:
  i. Declare an empty string that will hold decoded characters
  ii. Declare an alphabet string that holds all alphabetical characters
     ABCS := "abcdefghijklmnopqrstuvwxyz"
  iii. declare a map to hold indices of alphabetical letters
       map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4,
		      "f": 5, "g": 6, "h": 7, "i": 8, "j": 9,
		      "k": 10, "l": 11, "m": 12, "n": 13, "o": 14,
		      "o": 15, "q": 16, "r": 17, "s": 18, "t": 19,
		      "u": 20, "v": 21, "w": 22, "x": 23, "y": 24, "z": 25}
  iv. Iterate over the characters of the given string and on each iteration:
     > check if a lowercased version of the string character exists within the
        alphabet string
        > if it doesn't then add the string character as is to the decoded string
        > if it does, then declare a variable to establish if should uppercase the
            string or not  (string(char) == strings.ToUpper(string(char))
       >  find the index of the lowercased version of the alphabetical letter
       >  add 13 to the index to find the new proposed index of the decoded version
          of the character
          > check if the index is out of bounds by testing if it's greater than 25
             > if it is, then subtract 25 from the sum of the index and 13
             > use the result to subset the alphabet string to find the decoded character
             > ex: "v" has index of 21, 21 + 13 = 34, if 34 - 26 = 8, ALPHABET[8] = "i" as intended
       > if the flag for uppercasing the decoded character is set to true then uppercase
         the decoded char before pushing it to the string of decoded characters
   v. return the decoded string
*/

func main() {
	// Test cases:
	fmt.Println(decodeROT13String("EBG13 rknzcyr."))         //  "ROT13 example."
	fmt.Println(decodeROT13String("uryyb sebz bhgRefcnpr5")) // "hello from outErspace5"
}

func decodeROT13String(str string) string {
	decoded := ""

	for _, char := range str {
		strChar := string(char)
		upperCaseFlag := strChar == strings.ToUpper(strChar)

		unencrypted := decodeROT13Char(strChar)
		if upperCaseFlag {
			unencrypted = strings.ToUpper(unencrypted)
		}
		decoded += unencrypted
	}
	return decoded
}

func decodeROT13Char(str string) string {
	ABCS := "abcdefghijklmnopqrstuvwxyz"
	abcIndices := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4,
		"f": 5, "g": 6, "h": 7, "i": 8, "j": 9,
		"k": 10, "l": 11, "m": 12, "n": 13, "o": 14,
		"p": 15, "q": 16, "r": 17, "s": 18, "t": 19,
		"u": 20, "v": 21, "w": 22, "x": 23, "y": 24, "z": 25}

	idx, exists := abcIndices[strings.ToLower(str)]
	if !exists {
		return str
	} else {
		return string(ABCS[inBoundsIdx(idx)])
	}
}

func inBoundsIdx(idx int) int {
	newIdx := idx + 13
	if newIdx > 25 {
		newIdx -= 26
	}
	return newIdx
}
