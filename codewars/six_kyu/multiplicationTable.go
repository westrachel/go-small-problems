package main

import "fmt"

/* Write a function that takes an integer N and returns
an NxN multiplication table.

Algorithm:
i. declare an empty list that will contain lists of integers
ii. iterate from a rowNum of 1 to N and on each iteration:
    > declare a runningTotal variable initialized to rowNum
    > declare an empty row list that will contain integers
    > iterate from a colNum of 1 to N and on each iteration:
      > push the runningTotal variable to the row list
      > add the value assigned to rowNum to runningTotal
    > push the list created in the inner loop to the list
     created in step (i)
iii. return the list of lists

*/

func main() {
	// Test Cases:
	fmt.Println(multiplicationTable(0)) // []
	fmt.Println(multiplicationTable(1)) // [[1]]
	fmt.Println(multiplicationTable(3)) // [[1,2,3],[2,4,6],[3,6,9]]
	fmt.Println(multiplicationTable(5))
	/* [[1,2,3,4,5],[2,4,6,8,10],[3,6,9,12,15],[4,8,12,16,20],[5,10,15,20,25]] */
}

func multiplicationTable(size int) [][]int {
	var listOfLists [][]int

	for rowNum := 1; rowNum <= size; rowNum++ {
		runningTotal := rowNum
		var rowList []int
		for colNum := 1; colNum <= size; colNum++ {
			rowList = append(rowList, runningTotal)
			runningTotal += rowNum
		}
		listOfLists = append(listOfLists, rowList)
	}

	return listOfLists
}
