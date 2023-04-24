package main

import (
	"fmt"
	"strings"
)

/*
  Write a function that accepts an integer and builds
  a pyramid tower with the integer number of floors.

  Exs:
  tower(3)
  [
    "  *  ",
    " *** ",
    "*****"
  ]

  tower(6)
  [
    "     *     ",
    "    ***    ",
    "   *****   ",
    "  *******  ",
    " ********* ",
    "***********"
  ]

  Input: integer - represents the # of string elements
   in the returned list
  Output: list of string elements

  Star Pattern: N = 6
  > row 1: 1 star, 2N - 1 - 1 = 10 spaces
  > row 2: 3 stars, 2N - 1 - 3 = 8 spaces
  > row 3: 5 stars, 2N - 1 - 5 = 6 spaces
  > row 4: 7 stars, 2N - 1 - 7 = 4 spaces
  > row 5: 9 stars, 2N - 1 - 9 = 2 spaces
  > row 6: 11 stars = 2N - 1 stars, 0 spaces

  Algorithm:
  1. declare numStars and initialize it to 1
  2. declare totalChars and initialie it to 2N - 1, where N is the
     integer argument
  3. declare an empty list
  4. iterate from 0 to the given integer N minus 1 & on each iteration
     > create a string that concatenates:
        > (2N - 1 - numStars) spaces
        > numStars # of "*"
	> (2N - 1 - numStars) spaces
     > push this string to a list
     > add 2 to numStars
  5. return the list of string elements of stars/spaces
*/

func main() {
	fmt.Println(tower(3))
	fmt.Println(tower(6))
	fmt.Println(tower(15))
}

func tower(n int) []string {
	numStars, totalChars := 1, 2*n-1
	var tower []string
	for row := 0; row <= n-1; row++ {
		spaces := strings.Repeat(" ", (totalChars-numStars)/2)
		stars := strings.Repeat("*", numStars)
		tower = append(tower, fmt.Sprintf("%v%v%v", spaces, stars, spaces))
		numStars += 2
	}

	return tower
}
