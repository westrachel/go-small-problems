package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 Write a function that takes two lists of of strings a1 and a2 and
 returns a list of strings in a1 that are substrings of strings in a2.

  Reqs:
  > strings in a1 should be sorted lexicographically
  > case doesn't matter; return like cased words though
  > there should be no duplicates in the returned string

Algorithm:
  i. declare a string that will store all the words of the 2nd list
     and declare a map of addedWords to ensure that no duplicates are
     added to the list
 ii. iterate over the second list and:
       > push a lowercased version of each word to the string
       > after pushing each word push a 0 digit character to mark the
         start of a new word
 iii. declare an empty list that will hold the substrings
 iv. iterate over the first list and on each iteration:
     > lowercase the current string word
     > check if the lowercased word is contained within the mega string of
        all words from the 2nd list
     > if it is contained, then check if this lowercased string is already
       in the map of addedWords
        > if it's not in the map as a key, then don't push this string
          to substrings because it's already in substrings
 v. sort the list of substrings to ensure the words are ordered correctly
 vi. return the sorted list
*/

func main() {
	// Test cases:
	fmt.Println(inList([]string{"live", "arp", "strong"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"})) // ["arp", "live", "strong"]
	fmt.Println(inList([]string{"tarp", "mice", "bull"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"})) // []
	var empty []string
	fmt.Println(inList(empty, empty)) // []
	fmt.Println(inList([]string{"ect", "hE", "iNg", "ing", "omM", "oU", "tiOn"},
		[]string{"ect", "he", "ing", "ing", "omm", "ou", "tion"})) // ["ect", "hE", "iNg", "ing", "oU", "omM", "tiOn"]
}

func inList(s1, s2 []string) []string {
	words := containingString(s2)
	var substrings []string
	addedWords := make(map[string]bool)

	for _, str := range s1 {
		word := string(str)
		if strings.Contains(words, strings.ToLower(word)) {
			if !addedWords[word] {
				substrings = append(substrings, word)
				addedWords[word] = true
			}
		}
	}
	sort.Strings(substrings)
	return substrings
}

func containingString(list []string) string {
	allWords := ""
	for _, word := range list {
		lower := strings.ToLower(string(word))
		allWords += lower
		allWords += "0"
	}

	return allWords
}
