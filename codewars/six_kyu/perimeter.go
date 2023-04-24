package main

import "fmt"

/*
  Write a function that takes an integer n that represents the number
  of squares and returns the perimeter length of all the squares.

  Reqs:
  > the length of one side of each square follows fibonacci sequence
    > so if have 6 squares, their sides are of lengths:
         1, 1, 2, 3, 5, 8
  > the squares are numbered from 0, so if given n = 5, then there are
     actually 6 squares
      > in this case, need to find the first 6 numbers in the fibonacci
         sequence, add them together, and multiply the sum by 4

  Recursive Algorithm:
  > i. define a helper function that returns the sum of numbers in the
     fibonacci sequence
      > for a given integer n, that represents the amount of numbers from
         the fibonacci sequence that's desired:
          > return a list: [1, 1] that represents [numberInFibonacciSequence, sumOfFibNumbers] if n == 0
          > return a list: [1, 2] that represents [numberInFibonacciSequence, sumOfFibNumbers] if n == 1
          > return a list [numberInFibonacciSequence, sumOfFibNumbers] such that:
                numberInFibonacciSequence = fib(n-1)[0] + fib(n-2)[0]
		sumOfFibNumbers = fib(n-1)[1] + numberInFibonacciSequence
  > ii. find the number at the 1st index position in the list that represents the sum of individual
      lengths of each square and multiply it by 4 to calculate the perimeter

  Non-Recursive Algorithm:
  i. declare x1, x2, and sum variables assigned to 1, 1, and 2 respectively
  ii. iterate from 2 to the given integer n and on each iteration:
     <=> if n = 5, need to iterate 4 times to add 2, 3, 5, & 8 to the existing sum of 2
       > find the sum of x1 and x2
       > add the sum to the value assigned to the sum variable
       > reassign x2 to x1 and x1 to the sum
  iii. return the value assigned to sum multiplied by 4

*/

func main() {
	// Test Cases:
	fmt.Println(perimeter(5)) // 80 = 4 * (1 + 1 + 2 + 3 + 5 + 8)
	fmt.Println(perimeter(7)) // 216
}

/* Non-Recursion */
func perimeter(n int) int {
	x1, x2, sum := 1, 1, 2

	for idx := 2; idx <= n; idx++ {
		interimSum := x1 + x2
		sum += interimSum
		x2 = x1
		x1 = interimSum
	}

	return 4 * sum
}

/* Recursion:
func perimeter(n int) int {
	return 4 * fibonacci(n)[1]
}

func fibonacci(n int) []int {
	if n == 0 {
		return []int{1, 1}
	} else if n == 1 {
		return []int{1, 2}
	} else {
		priorNumSum := fibonacci(n - 1)
		newNum := fibonacci(n - 1)[0] + fibonacci(n - 2)[0]
		newSum := newNum + priorNumSum[1]
		return []int{newNum, newSum}
	}
}
*/
