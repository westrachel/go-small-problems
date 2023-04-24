package main

import (
	"fmt"
	"math"
)

/* Write a function that takes an amount of money and returns
   a list representing the smallest number of bills needed
   to represent the amount of money in change

The bills in the returned list need to be output in a specific order:
    element @ slice[0] represents $1 bills
    element @ slice[1] represents $5 bills
    element @ slice[2] represents $10 bills
    element @ slice[3] represents $20 bills
    element @ slice[4] represents $50 bills
    element @ slice[5] represents $100 bills

 Reqs:
 > dollar amount inputs are > 0 and < 1000

Algorithm:
 i. declare a map to store the counts associated with each dollar bill type
    { 100.0: 0, 50.0: 0, 20.0: 0, 10.0:0, 5.0:0, 1.0: 0}
 ii. declare a dollarAmount variable that is assigned initially to the float64
      value of the given dollar amount argument
 iii. start a loop that will continue to iterate until there's no remaining
      change and on each iteration, find the next largest float64 divisor that
      results in a quotient > 1
         > bill divisor can be found by determining if the remaining change
            left is >= 100.0, >= 50.0, >= 20.0, >= 10.0, >= 5.0, or >= 1.0
         > find the quotient from dividing the remaining dollar amount by the
            bill divisor
         > round the quotient down to the nearest integer
         > reassign the bill key's corresponding value to the rounded quotient
         > reassign the remaining dollar amount to the result of subtracting
             the product of the bill by rounded quotient from the current
              value assigned to the remaining dollar amount
iv. return a slice that contains the following ordered values:
      map[1], map[5], map[10], map[20], map[50], map[100] */

func main() {
	// Test cases:
	fmt.Println(giveChange(365)) // [0,1,1,0,1,3]
	fmt.Println(giveChange(217)) // [2,1,1,0,0,2]
}

func giveChange(amount int) []float64 {
	change := map[float64]float64{100.0: 0.0, 50.0: 0.0, 20.0: 0.0, 10.0: 0.0, 5.0: 0.0, 1.0: 0.0}
	remainingDollarAmt := float64(amount)

	for remainingDollarAmt > 0 {
		bill := nextBillAmt(remainingDollarAmt)
		count := math.Floor(remainingDollarAmt / bill)
		change[bill] = count
		remainingDollarAmt -= (bill * count)
	}

	return []float64{change[1.0], change[5.0], change[10.0], change[20.0], change[50.0], change[100.0]}
}

func nextBillAmt(remainingDollarAmt float64) float64 {
	if remainingDollarAmt >= 100.0 {
		return 100.0
	} else if remainingDollarAmt >= 50.0 {
		return 50.0
	} else if remainingDollarAmt >= 20.0 {
		return 20.0
	} else if remainingDollarAmt >= 10.0 {
		return 10.0
	} else if remainingDollarAmt >= 5.0 {
		return 5.0
	} else if remainingDollarAmt >= 1.0 {
		return 1.0
	} else {
		return 0.0
	}
}
