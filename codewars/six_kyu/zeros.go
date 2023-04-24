package main

import (
	"fmt"
)

/* Write a function that calculates the number of trailing zeros in a factorial

Algorithm:
  i. declare a product variable to 1
  ii. iterate backwards from the given number to 1
  iii. multiply the product by the current number being iterated over
  iv. convert the product to a string
  v. declare a zeroCount variable that's initialized to 0
  vi. check if the last character of the string product is 0
       > if it's not then immediately return zeroCount
       > if there is a trailing zero, then iterate backwards over
          the string product and on each iteration add 1 to the
          zeroCount as long as the current string digit being iterated
          over is equal to 0
          > if encounter a nonzero string digit then break out of the
            loop
  vii. return zeroCount
*/

func main() {
	// Test cases:
	fmt.Println(zeros(6))  //  6! = 720 --> 1
	fmt.Println(zeros(12)) //  12! = 479001600 --> 2
}

func zeros(num int) int {
	product := stringFactorial(num)
	maxIdx := len(product) - 1
	lastDigit := string(product[maxIdx])

	if lastDigit != "0" {
		return 0
	}

	return countLastZeros(product)
}

func stringFactorial(num int) string {
	product := 1
	for factor := num; factor > 1; factor-- {
		product *= factor
	}
	return fmt.Sprintf("%v", product)
}

func countLastZeros(str string) int {
	zeroCount := 0
	for idx := len(str) - 1; idx >= 0; idx-- {
		digit := string(str[idx])

		if digit != "0" {
			break
		} else {
			zeroCount += 1
		}
	}

	return zeroCount
}
