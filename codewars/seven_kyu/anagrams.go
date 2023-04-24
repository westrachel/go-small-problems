package main

import (
	"fmt"
	"strings"
)

// Write a function that accepts 2 strings and determines if they are
// anagrams
// Reqs:
// > anagrams are case insensitive, but must have same length
//     <=> ie
// > return a boolean

// Algorithm:
// i. check if the strings have the same length; if they don't then return
//  false
// ii. declare variables that point to downcased versions of the provided
//  strings to ensure letter comparisons are of the same case
// iii. iterate through each of the given strings & create two maps that
//   whose keys represent each character of each string and whose values
//   keep track of the counts of that character in its respective string
// iv. iterate through one of the strings characters again and check if
//   its counts are equivalent to its counts in the other string
//    > if the counts are different reassign then return false
// v. return true if none of the prior steps resulted in stopping the
// function's execution and returning false

// Time complexity of worst case simplifies to O(N)
// rationale:
// > have to iterate through each of the strings once to create the maps
//   > b/c return false right away if strings have different lengths
//    know that the strings must have the same length once get to the create
//    maps section (so each string has N number of characters)
//    <=> so, the iterations underlying creating the maps account for 2N
//       steps
// > then iterate through one of the strings again to compare counts of that
//    character in each of the given strings based on map lookups of the character
//     <=> have to iterate through N characters, so this adds N steps
// > reading from a map w/in each iteration is O(1), so comparing a character's
//    counts in each string doesn't significantly add to the time complexity

// Altering the algorithm to sort the strings would not provide a time benefit
// b/c sorting may take O(N log N) steps, which for a large number of N elements
// exceeds N

// Space Complexity in worst case = O(N)
// rationale: create 2 new maps that store at most N unique key/value pairs

func main() {
	// Test Cases:
	fmt.Println(anagrams("foefet", "toffee"))         // true
	fmt.Println(anagrams("madame", "adam"))           // false
	fmt.Println(anagrams("Buckethead", "DeathCubeK")) // true
	fmt.Println(anagrams("sun", "could"))             // false
	fmt.Println(anagrams("would", "could"))           // false
}

func anagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	s1, s2 := strings.ToLower(str1), strings.ToLower(str2)
	counts1, counts2 := charCounts(s1), charCounts(s2)

	for i, _ := range s1 {
		char := string(s1[i])
		if counts1[char] != counts2[char] {
			return false
		}
	}
	return true
}

func charCounts(str string) map[string]int {
	counts := make(map[string]int)
	for idx, _ := range str {
		char := string(str[idx])
		_, keyExists := counts[char]
		if keyExists {
			counts[char] += 1
		} else {
			counts[char] = 1
		}
	}
	return counts
}
