package main

import (
	"fmt"
	"strings"
)

/* Write a function that takes a string of words and returns the highest
   scoring word.

  Reqs:
  > Each letter's points is based on its location in the alphabet a = 1, b = 2, c = 3
  > if multiple words have the same score, return the one that appears earlier in the string
  > all letters are lowercase
  > words are separated by single spaces and consist only of alphabetical words

Algorithm:
  i. declare a map where each key is a lowercased string letter in the alphabet and the value
   is the associated point amount
  ii. split the string into words based on splitting at the space delimiter
  iii. declare a maxWord variable that's initialized to an empty string
  iv. declare a maxScore variable that's initialized to 0
  v. iterate over the words and on each iteration:
      > declare a sum variable that's intialized to 0
      > iterate over the characters in the string and on each iteration:
          > find the value associated with the character in the map
          > add the value to the sum
      > check if the sum is greater than the value assigned to maxScore
         > if it is, then reassign maxScore to the current sum and
            reassign the maxWord variable to the current word being iterated over
  vi. return the value assigned to the maxWord variable
*/

func main() {
	// Test cases:
	fmt.Println(highestScore("apple banana pear")) // "apple"
	fmt.Println(highestScore("hello world"))       // "world"
}

func highestScore(str string) string {
	words := strings.Split(str, " ")
	maxWord := ""
	maxScore := 0

	for _, word := range words {
		sum := 0
		for _, char := range string(word) {
			sum += charScore(string(char))
		}
		if sum > maxScore {
			maxScore = sum
			maxWord = string(word)
		}
	}

	return maxWord
}

func charScore(letter string) int {
	scores := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7,
		"h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14,
		"o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21,
		"v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	}
	return scores[letter]
}
