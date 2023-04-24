package main

import (
	"fmt"
	"strings"
)

// Write a function that accept a string and removes the egg encoding
// Reqs: 2
// > "egg" encoding entails an "egg" being inserted directly after
//   each consonant
// > vowels = "AEIOUY"
// > can assume string has at least one character & only contains
//   alphabetical letters (ie no instances of 5, -, !, etc)

// Algorithm:
// i. declare a variable to store characters that aren't part of any
//   "egg" encoding
// ii. declare a starting index variable that points to 0
// iii. declare a vowels constant string that contains all vowels
// iv. start a for loop and on each iteration:
//   > Add the character to the string varaible that's keeping track
//     of non-egg string characters
//   > check if the character at the index position in the given
//     string is a constant (by checking if the vowels string
//     contains an uppercased version of the character)
//   > if the character is a constant then add 4 to the number
//     that the index variable points to
//   	> if the character is a vowel, then add 1 to the number
//       that the index variable points to
// v. return the string created w/o the eggs

func main() {
	// Test Cases:
	fmt.Println(unscrambleEggs("Beggegeggineggneggeregg")) // "Beginner"
	fmt.Println(unscrambleEggs("yeLeggleggOwegg"))         // "yeLlOw"
	fmt.Println(unscrambleEggs("a"))                       // "a"
	fmt.Println(unscrambleEggs("segguneggsegghegginegge")) // "sunshine"
}

func unscrambleEggs(str string) string {
	const Vowels = "AEOIUY"
	var idx int
	var unencoded string
	for {
		char := string(str[idx])
		unencoded += char
		if strings.Contains(Vowels, strings.ToUpper(char)) {
			idx += 1
		} else {
			idx += 4
		}
		if idx >= len(str) {
			break
		}
	}

	return unencoded
}
