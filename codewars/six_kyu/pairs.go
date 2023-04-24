package main

import "fmt"

/* Write a function that splits a given string into pairs of
two characters. If there's an odd number of characters, then
replace the missing second character of the final pair with
an underscore ('_').

 Reqs:
   > should return a list of pairs
   > return empty list if given an empty string

Algorithm:
i. declare an empty list
ii. declare a variable that's initialized to the length of the
    string; this variable keeps track of whether or not the ending
    index of a slice is out of bounds for the given list
      slice[2:4]  <=> 4 is not the ending index of the slice, 3 is,
        which is why this variable is initialized to the string's
        length, which is 1 greater than the largest index in the list
iii. iterate from 0 to the length of the string - 1 and on each
    iteration:
      > find the ending index of the slice, which is the current
        index plus 2
      > check if the ending index of the slice is greater than
         the out of bounds index
          > if it is, then find the value of the string at the
            current index and add an underscore of the string
             > push the resulting string to the list of pairs
          > if it's not, then find the slice of the string and
             push this slice to the list of pairs
iv. return the list of pairs

   // splits:
   str := 'abcdef'
     str[0:2] = 'ab'
     str[2:4] = 'cd'
     str[4:6] = 'ef'
   str := 'abcde'
     str[0:2] = 'ab'
     str[2:4] = 'cd'
     str[4:6] = 'ef'
*/

func main() {
	// Test cases:
	fmt.Println(pairs("abc"))    // ['ab', 'c_']
	fmt.Println(pairs("abcdef")) // ['ab', 'cd', 'ef']
	fmt.Println(pairs("abcde"))  // ['ab', 'cd', 'e_']
	fmt.Println(pairs(""))       // []
}

func pairs(str string) []string {
	var pairs []string
	outOfBounds := len(str)

	for idx := 0; idx <= len(str)-1; idx += 2 {
		endIdx := idx + 2
		if endIdx > outOfBounds {
			pair := string(str[idx]) + "_"
			pairs = append(pairs, pair)
		} else {
			pairs = append(pairs, string(str[idx:endIdx]))
		}
	}

	return pairs

}
