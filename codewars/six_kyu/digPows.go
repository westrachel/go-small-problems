package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
  Write a function that takes an integer n and a power
  and returns -1 or k such that k is an integer that
  multiplied by the given n results in a product such that:
   a^power + b^(power + 1) + c^(power + 2) + d^(power + 3) = abcd(k)
  	where n = abcd	<=> a, b, c, d represent individual
	 digits of n

  Algorithm:
  i. declare a total variable initialized to 0
  ii. convert the given integer n into a string
  iii. iterate over the characters in the given string; on each iteration:
      > convert the string digit to a number
      > raise the number to the value assigned to power
      > add the result to the value assigned to total
      > add 1 to the value assigned to power
  iv. divide the total value by the integer n and check if the remainder is 0
      > if the remainder is zero then return this quotient
      > if the remainder is not zero then return -1
*/

func main() {
	// Test cases:
	fmt.Println(digPow(89, 1))    // 1 	<=> b/c 8¹ + 9² = 89 = 89 * 1
	fmt.Println(digPow(92, 1))    // -1 	<=> b/c no k such as 9¹ + 2² equals 92 * k
	fmt.Println(digPow(695, 2))   // 2      <=> b/c 6² + 9³ + 5⁴= 1390 = 695 * 2
	fmt.Println(digPow(46288, 3)) // 51     <=> b/c 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
}

func digPow(n, pow int) int {
	total := sumPows(n, pow)
	quotient := total / n

	if total%n == 0 {
		return quotient
	} else {
		return -1
	}
}

func sumPows(n, pow int) int {
	total := 0
	strN := fmt.Sprintf("%v", n)

	for _, char := range strN {
		digit, _ := strconv.Atoi(string(char))
		result := math.Pow(float64(digit), float64(pow))
		total += int(result)
		pow += 1
	}
	return total
}
