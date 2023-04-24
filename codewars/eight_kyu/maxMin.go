package main

import "fmt"

// write a function that returns the max and min of a list of integers
// [4,6,2,1,9,63,-134,566]             -> max = 566, min = -134
// [-52, 56, 30, 29, -54, 0, -110, 24] -> min = -110, max = 56
// [42, 54, 65, 87, 0, 33, 29, 19, 11] -> min = 0, max = 87
// [5]                                 -> min = 5, max = 5

func main() {
	var arr1 [1]int = [1]int{5}
	var s []int = arr1[:]
	f1 := maxMin
	result1, result2 := f1(s)
	fmt.Printf("Max, Min: ", result1, result2)
}

func maxMin(list []int) (max, min int) {
	max, min = list[0], list[0]
	for i := 0; i < len(list); i++ {
		num := list[i]
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return max, min
}
