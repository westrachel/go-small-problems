package main

import (
	"fmt"
	"strings"
)

/*
  Write a function that takes a string of words separated by spaces
  and reverses the order of the words and then for each word in the
  string all the characters are reversed except for the last character.
  That should remain at the end as it represents punctuation unless
  the string is empty.

  Algorithm:
   i. declare reversedWords and initialize it to an empty string
   ii. check if the given string is empty, in which case return reversedWords
       immediately
   iii. declare and initialize strings that point to the punctuation and all
         the non-punctuation characters in the given string
          > puncutation is always the last character in the given string
          > non-puncutaton ranges from the 0th index to the character before
            the last character in the string
   iv. split the non-punctuation string into separate words
   v. start a loop that will iterate backwards over the slice of words
      > use an index as part of the condition whose initial value is
         len(words) - 1 to 0 and on each iteration:
          > find the word being iterated over <=> it's the element at the
              current index position in the slice of words
          > split the word into its characters
          > assign to reversedString an empty string
          > iterate over the characters of the word backwards from the
            last index (len(chars) - 1)
          > on each iteration push each character to reversedString
      > after iterating, add the last element, the punctuation, to
         reversedString
            punctuation := chars[len(chars) - 1]
      > add the reversedWord to reversedWords
   vi. return reversedWords
*/

func main() {
	// Test cases:
	fmt.Println(esrever("hello world."))                   // "dlrow olleh."
	fmt.Println(esrever("Much l33t?"))                     // "t33l hcuM?"
	fmt.Println(esrever("goodbye M00n see you tomorrow!")) // worromot uoy ees n00M eybdoog!
	fmt.Println(esrever("tacocat!"))                       // "tacocat!"
	fmt.Println(esrever(""))                               // ""
}

func esrever(words string) string {
	reversedWords := ""
	if len(words) == 0 {
		return reversedWords
	}
	punctuation, onlyChars := words[len(words)-1], words[:len(words)-1]
	allWords := strings.Split(onlyChars, " ")

	for idx := len(allWords) - 1; idx >= 0; idx-- {
		word := allWords[idx]
		chars := strings.Split(word, "")
		reversedWord := ""

		for i := len(chars) - 1; i >= 0; i-- {
			reversedWord += chars[i]
		}
		if idx != 0 {
			reversedWord += " "
		}
		reversedWords += reversedWord
	}

	return reversedWords + string(punctuation)
}
