package main

import (
	"fmt"
	"strings"
)

// Write a function that returns a boolean indicating if a given string
// contains the same number of Xs and Os

// Algorithm:
// i. declare a map that's initialized to have keys for uppercased
//  x's and o's whose values are all zero
// ii. iterate through the given string and check if the uppercased version
//     of each string character is a key in the map
//     > if it is a key then add one to the corresponding value
//     > if isn't a key then do nothing
// iii. return true if the value associated with the 'X' key is equal to the
//    value associated with the 'O' key

func main() {
	// Test cases:
	fmt.Println(XO("ooxx"))    // true
	fmt.Println(XO("xooxx"))   // false
	fmt.Println(XO("ooxXm"))   // true
	fmt.Println(XO("zpzpzpp")) // true
	fmt.Println(XO("zzoo"))    // false
}

func XO(str string) bool {
	counts := map[string]int{"X": 0, "O": 0}
	for _, char := range str {
		strChar := strings.ToUpper(string(char))
		_, keyExists := counts[strChar]
		if keyExists {
			counts[strChar] += 1
		}
	}
	fmt.Println(counts)
	return counts["X"] == counts["O"]
}
