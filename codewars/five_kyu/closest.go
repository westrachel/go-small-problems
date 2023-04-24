package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
 Write a function that takes a string of integers and returns 2 slices
 that contain the weight of 2 integers in the original string, their
 index position, and their integer value.

  Input:
   > string of unknown length
   > contains integers that are separated by single spaces

  Output:
  > a slice of 2 slices that contain:
   > the weight of the integer = sum of individual digits
       99's weight = 18
      100's weight = 1
   > the index position
       > this is not the original string index position
       > turn the string into an array by splitting at spaces
       > the index position of the integer in ther array is the index
         that should be returned
   > the integer

 Requirements:
  > the 2 integers selected from the original string should have the
    smallest difference between their weights
  > if there are 2 different sets of integers with the same smallest
     difference, then choose to return the pair of integers that are
     smaller
     > ie if 2 and 4 are weights of 2 numbers and 4 and 6 are another
        pair with the same difference of 2, then return the pair
        of numbers that have weights 2 and 4, b/c those weights are
        smaller
  > if all integers in the string have the same weight, then return
     the integers w/ the smallest indices
  > number of integers in the original string is 0 or >= 2
     > n = 0 looks like an empty string (not a string of multiple spaces)
        > for n = 0 return an empty slice
     > won't be given non-integer words in the string

  Algorithm:
  i. if the length of the given string is 0 then return an empty slice
  ii. split the given integer of strings at spaces
  iii. declare a weights slice that's initially empty and a map that will store weights
       as keys and values as slices that contain the weight's initial integer value
       and index position in the array of string integers
  iv. iterate over the strings of integers and on each iteration:
         > iterate over each individual string and find the sum of its digits
         > push the sum of digits (the weight) to the weights slice
         > check if the sum of digits (the weight) already exists as a key in the
           map
            > if it does then add to the closest slice 2 slices that contain the slices:
              [weight, map[weight][0], map[weight][1]]
	      [weight, currentIndex, strconv.AtoI(currentStrInteger)]
         > assign the sum of its digits as a key in the map and the value as a
              slice containing the index position of the string integer in the
              array and an integer value of the string integer
  v. sort the weights slice from smallest to largest
  vi. declare a minDifference value that's equal to the difference of the weights at index
       0 and index 1 of the sorted weights array
  vii. declare w1 and w2 variables that are initialized to the weights at the 0th and 1st
       index of the sorted weights array
  viii. iterate from an index of 1 to an index of len(weights) - 2 and on each iteration:
       > find the difference between the weights at the current index and the current index
         plus 1 position of the sorted weights array
       > check if the difference is smaller than the current mindDifference
         > if it is, the reassign minDifference to the current difference and reassign
            w1 and w2 to the weights at the current index and current index + 1 position
            in the sorted weights array
  ix. return a slice that contains the 2 slices:
              [w1, map[w1][0], map[w1][1]]
	      [w2, map[w2][0], map[w2][1]]
*/

func main() {
	fmt.Println(closest(""))                                                 //   []
	fmt.Println(closest("103 123 4444 99 2000"))                             //   [[2, 4, 2000], [4, 0, 103]]
	fmt.Println(closest("80 71 62 53"))                                      //   [[8, 0, 80], [8, 1, 71]]
	fmt.Println(closest("444 2000 445 544"))                                 //   [[13, 2, 445], [13, 3, 544]]
	fmt.Println(closest("444 2000 445 644 2001 1002"))                       //   [[3, 4, 2001], [3, 5, 1002]]
	fmt.Println(closest("239382 162 254765 182 485944 468751 49780 108 54")) // [[9, 1, 162], [9, 7, 108]]
	fmt.Println(closest("54 239382 162 254765 182 485944 468751 49780 108")) // [[9, 0, 54], [9, 2, 162]]
}

func closest(str string) [][]int {
	if len(str) == 0 {
		var empty [][]int
		return empty
	}
	strNums := strings.Split(str, " ")

	var weights []int
	weightVals := make(map[int][]int)

	for idx, chars := range strNums {
		strNum := string(chars)
		weight := calcWeight(strNum)
		weights = append(weights, weight)
		valSlice, keyExists := weightVals[weight]
		num, _ := strconv.Atoi(strNum)

		if keyExists {
			return [][]int{[]int{weight, valSlice[0], valSlice[1]}, []int{weight, idx, num}}
		} else {
			weightVals[weight] = []int{idx, num}
		}
	}
	sort.Ints(weights)
	return minDifferencePairs(weights, weightVals)
}

func minDifferencePairs(weights []int, weightVals map[int][]int) [][]int {
	w1, w2 := weights[0], weights[1]
	minDiff := w2 - w1
	for idx := 1; idx < len(weights)-2; idx++ {
		newW1, newW2 := weights[idx], weights[idx+1]
		currDiff := newW2 - newW1
		if currDiff < minDiff {
			minDiff, w1, w2 = currDiff, newW1, newW2
		}
	}

	s1, s2 := findSubslice(w1, weightVals), findSubslice(w2, weightVals)
	return [][]int{s1, s2}
}

func findSubslice(key int, valuesMap map[int][]int) []int {
	vals := valuesMap[key]
	return []int{key, vals[0], vals[1]}
}

func calcWeight(strNum string) int {
	sum := 0
	for _, digit := range strNum {
		num, _ := strconv.Atoi(string(digit))
		sum += num
	}

	return sum
}
