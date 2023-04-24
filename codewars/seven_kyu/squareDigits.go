package main

import (
	"fmt"
	"math"
	"strconv"
)

// Write a function takes an integer and returns a concatenated integer
// version of the integer where each digit within the given number is
//  squared. Can assume only positive integer inputs.

// Ex: input 9119 -> return 811181
//      9^2 is 81, 1^2 is 1

// Algorithm:
// i. convert the given integer to a string
// ii. iterate over the string's characters
// iii. on each iteration convert the string to a number using previously
//   defined toNumber() function
// iv. square the number
// v. convert the squared number to a string
// vi. add the string digit to a variable string keeping track of all squared
//    digits
// vii. convert the squared digit string into a number & return this number

func main() {
	// Test cases:
	fmt.Println(squareDigits(9119) == 811181)
	fmt.Println(squareDigits(24) == 416)
	fmt.Println(squareDigits(139) == 1981)
}

func squareDigits(num int) int {
	digits := strconv.Itoa(num)
	var squared string

	for i, _ := range digits {
		num := toNumber(string(digits[i]))
		squared += strconv.Itoa(num * num)
	}

	return toNumber(squared)
}

func toNumber(str string) int {
	multiplier := 1
	var number int

	if str[0:1] == "-" {
		multiplier *= -1
		number = charsAsDigits(str[1:])
	} else {
		number = charsAsDigits(str)
	}
	return multiplier * number
}

func charsAsDigits(str string) int {
	var strDigits = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	num := 0
	for i, exponent := 0, len(str)-1; i < len(str); i, exponent = i+1, exponent-1 {
		digit := strDigits[str[i:i+1]]
		num += int(math.Pow(10, float64(exponent))) * digit
	}
	return num
}
