package main

import (
	"fmt"
	"math"
)

// Write a function that can transform a string into an integer.
// Reqs:
//  > inputs can be positive or negative string representations of a number
//  > inputs are integer representations (so don't have a nonsensical letter
//     included like 'b')

func main() {
	// Test Cases:
	fmt.Println(toNumber("1234"))    // 1234
	fmt.Println(toNumber("605"))     // 605
	fmt.Println(toNumber("-1405"))   // -1405
	fmt.Println(toNumber("-7"))      // -7
	fmt.Println(toNumber("-239670")) // -239670
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
