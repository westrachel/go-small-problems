package main

import "fmt"

/*
  Write a function that takes an nxn slice of sublices of elements
  and returns a slice that contains all elements of the given elements
  in snail order.

  Input: nxn slice of subslices of integers
  Output: slice of ordered integers

  Reqs:
  > if given 0x0 slice of subslices return a flattened empty slice

  Algorithm:
  i. declare an empty slice to store the ordered elements
  ii. iterate over the slice of subslices and on each iteration:
      > push all elements from the first row to the new ordered slice
      > reslice the given slices to include all subslices from index 1 to the end
         of the subslices
      > push the last element of all remaining sublices to the new slice
         > need to iterate over each of the remaining sublices
         > on each iteration add the row's last element (element at index = len(row) -1)
            to the ordered slice
         > reslice the row to contain elements from the 0th index to 2nd to last index
      > push in reverse order the elements in the last row to the new slice
      > reslice the subslices to exclude the last row
      > push the 0th indexed element from all the remaining rows to the ordered slice of all elements
  iii. repeat step 2 working inwards
      > repeat by deleting elements from each array when push them to the new slice
*/

func main() {
	// Test cases:
	var emptySlices [][]int
	fmt.Println(snail(emptySlices))                                             // []
	fmt.Println(snail([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})) // []int{1,2,3,6,9,8,7,4,5}
	fmt.Println(snail([][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}})) // []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(snail([][]int{[]int{1, 2, 3, 1}, []int{4, 5, 6, 4}, []int{7, 8, 9, 7}, []int{7, 8, 9, 7}}))
	// []int{1,2,3,1,4,7,7,9,8,7,7,4,5,6,8,9}
}

func snail(slices [][]int) []int {
	var ordered []int
	for len(slices) > 0 {
		ordered = transferRow(ordered, slices[0])
		slices = slices[1:]
		for rowIdx := 0; rowIdx < len(slices); rowIdx++ {
			row := slices[rowIdx]
			lastIdx := len(row) - 1
			ordered = append(ordered, row[lastIdx])
			slices[rowIdx] = row[:lastIdx]
		}
		if len(slices) > 0 {
			ordered = transferRow(ordered, reverse(slices[len(slices)-1]))
			slices = slices[:len(slices)-1]
			for rowIdx := len(slices) - 1; rowIdx >= 0; rowIdx-- {
				row := slices[rowIdx]
				ordered = append(ordered, row[0])
				slices[rowIdx] = row[1:]
			}
		}
	}

	return ordered
}

func transferRow(newList, list []int) []int {
	for _, element := range list {
		newList = append(newList, element)
	}

	return newList
}

func reverse(slice []int) []int {
	var reversed []int

	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}
