package main

import (
	"fmt"
	"strings"
)

/* Write a function that returns a dollar amount of
 all the loose change found.

Input:
 > string containing words describing the loose
     change
 > possible change that can be collected:
    > penny: 0.01
    > nickel: 0.05
    > dime: 0.10
    > quarter: 0.25
    > dollar: 1.00
 Output:
    > string representation of a dollar amount out to
      the penny

Algorithm:
i. declare a map that stores each string change name and its
    corresponding float change amount
ii. split the given string of loose change at each space
iii. declare a total change variable that's assigned to 0
iv. iterate over the split string and:
    > look up the change amount's corresponding value in the map
    > add that value to the running total
v. return the floating dollar amount as part of a string that's
   prepended with the dollar sign */

func main() {
	// Test cases:
	fmt.Println(totalChange("nickel penny dime dollar"))                                                    // "$1.16"
	fmt.Println(totalChange("dollar dollar quarter dime dime"))                                             // "$2.45"
	fmt.Println(totalChange("penny"))                                                                       // "$0.01"
	fmt.Println(totalChange("dime"))                                                                        // "$0.10"
	fmt.Println(totalChange("dollar dollar dollar dollar dollar dollar dollar dollar dollar dollar penny")) // "$10.01"
}

func totalChange(change string) string {
	changeAmts := map[string]float64{"penny": 0.01, "nickel": 0.05, "dime": 0.10, "quarter": 0.25, "dollar": 1.00}
	looseChange := strings.Split(change, " ")
	total := 0.0

	for _, billCoin := range looseChange {
		total += changeAmts[billCoin]
	}
	return fmt.Sprintf("$%.2f", total)
}
