package main

import (
	"fmt"
	"strings"
)

/*
  Write a function that takes a string and converts it into a hastag

  Reqs:
  > if new string to return is >140 chars long then return false
  > capitalize each word in the returned string
  > remove any spaces
  > if input is empty or what you'd return w/o the hashtag is empty
     then return false
  > return needs to start with a hashtag

  Algorithm:
  1. check if the length of the string is 0; if so, then return "false"
  2. declare a smushed string to push characters to
  3. iterate over the indices and characters of the given string and on each
    iteration:
     > check if the character is a space
        > if it's not, then check if character at the index prior to the
           current index is a space
           > if it's not, then push the current non-space character to
             smushed string
           > if the prior character is a space, then capitalize the character
             and push it to the smushed string
  4. check if the smushed string's length is 0; if it is then return "false"
      > otherwise, return the result of concatenating a "#" with smushed string
*/

func main() {
	// Test cases:
	fmt.Println(hashTag(" Hello there thanks for stopping by")) //  "#HelloThereThanksForStoppingBy"
	fmt.Println(hashTag("    Hello     World   "))              // "#HelloWorld"
	fmt.Println(hashTag(""))                                    // false
	fmt.Println(hashTag("             "))                       // false
}

func hashTag(str string) string {
	if len(str) == 0 {
		return "false"
	}

	content := smush(str)
	if len(content) == 0 {
		return "false"
	}
	return "#" + content
}

func smush(str string) string {
	smushed := ""

	for idx, char := range str {
		letter := string(char)
		if letter != " " {
			if idx == 0 || string(str[idx-1]) == " " {
				letter = strings.ToUpper(letter)
			}
			smushed += letter
		}
	}
	return smushed
}
