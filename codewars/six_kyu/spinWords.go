package main

import (
	"fmt"
	"strings"
)

/*
Write a function that takes in a string of one or more words, and returns the string
 with words that consist of >= five letters reversed

  Reqs:
  > spaces delimit words if the string has >1 word

  Algorithm:
  i. split the string into words by splitting at spaces
  ii. declare a new empty stirng that will hold reversed/non-reversed string words
  iii. iterate over the words and on each iteration:
      > check if the length of the word is >= 5
      > if it's not, then push the word as is the string
      > if it is, then reverse the characters in the word and push the reversed
        word to the string
      > if the index of the current iteration is not equal to the last index in the
        list of words, then push a space character to the string

  Reverse:
   i. declare reverse variable that's intialized to an empty string
   ii. iterate backwards over the indices of the given string
   iii. on each iteraton, push the character at the index position in the string
        to the string assigned to reverse
   iv. return reverse
*/

func main() {
	// Test Cases:
	fmt.Println(spinWords("Hey fellow warriors"))  //"Hey wollef sroirraw"
	fmt.Println(spinWords("This is a test"))       // "This is a test"
	fmt.Println(spinWords("These"))                // "esehT"
	fmt.Println(spinWords("This is another test")) // "This is rehtona test"
}

func spinWords(str string) string {
	words, reversed := strings.Split(str, " "), ""

	for idx, chars := range words {
		word := string(chars)
		if len(word) >= 5 {
			reversed += reverse(word)
		} else {
			reversed += word
		}
		if idx != len(words)-1 {
			reversed += " "
		}
	}
	return reversed
}

func reverse(str string) string {
	var reversed string
	for idx := len(str) - 1; idx >= 0; idx-- {
		reversed += string(str[idx])
	}
	return reversed
}
