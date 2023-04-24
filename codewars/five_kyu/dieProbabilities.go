package main

import (
	"fmt"
	"math"
)

/*
  Write a function that takes a target sum and a number of die
  and returns the probability of rolling numbers that sum to
  the target.

  Reqs:
   > don't need to round the probabilities
   > dies have 6 sides

  Example: inputs = (8, 3)
  21 combinatios sum to 8:
  {2,3,3}, {3,2,3}, {3,3,2}
  {2,2,4}, {2,4,2}, {4,2,2}
  {1,3,4}, {1,4,3}, {3,1,4}, {4,1,3}, {3,4,1}, {4,3,1}
  {1,2,5}, {1,5,2}, {2,1,5}, {5,1,2}, {2,5,1}, {5,2,1}
  {1,1,6}, {1,6,1}, {6,1,1}

  probability = 21/(6*6*6) = 0.09722222222222222

  Algorithm:
  i. find all the permutations of length N that can be
     built with the numbers 1-6;  N = number of dies
     > define a function that:
         > declare empty slices of sublices of integers to store all
          combos and currentSlices
         > initialize these arrays to contain the numbers 1 -> 6
         > find a slice that contains numbers 2 through N inclusive
             > this will represent iterationNumbers
             > the idea is to create subslices whose lengths match
                the iterationNumbers and that can be added to over
                time to create all the different permutations of
                the different die side numbers (which range from 1-6)
         > iterate over the slice of iterationNumbers and on each iteration:
            > declare an empty slice of subslices of integers nextSubslices
            > iterate over the currentSlices and on each iteration:
               > iterate over the die side numbers: 1 -> 6
               > add the die side numbers to the subslice of currentSlice
                 being iterated over
               > push the subslice created to nextSublices and allSubslices
         > after iterating over all the subslices in currSlices, reassign
            currSlices to nextSubslices so that combinations grow in length
      > return all permutations found
  ii. filter through all the permutations and select only the permutations
       that add to the target and have length N
  iii. count the number of permutations that sum to the target
       <=> assign this value to a numerator
  iv. return the numerator divided by the result of calculating
      6 to the power of N

*/

func main() {
	fmt.Println(RolldiceSumProb(8, 3))  // 0.09722222222222222
	fmt.Println(RolldiceSumProb(11, 2)) // 0.0555555555
	fmt.Println(RolldiceSumProb(8, 2))  //  0.13888888889
	fmt.Println(RolldiceSumProb(4, 2))  // 0.0833333333 <=> 3/36
	fmt.Println(RolldiceSumProb(22, 3)) // 0
}

func RolldiceSumProb(targetSum, numDice int) float64 {
	denominator := math.Pow(float64(6), float64(numDice))
	combos := permutations(numDice, []int{1, 2, 3, 4, 5, 6})
	successCombos := filterByLengthSum(targetSum, numDice, combos)

	return float64(len(successCombos)) / denominator
}

func permutations(length int, elements []int) [][]int {
	all, currSlices := diceSlices(), diceSlices()
	iterationNums := rangeNums(2, length)
	for idx, _ := range iterationNums {
		fmt.Println("idx: ", idx)
		var nextSlices [][]int
		for _, subslice := range currSlices {
			for _, num := range elements {
				newSubslice := append(subslice, num)
				all = append(all, newSubslice)
				nextSlices = append(nextSlices, newSubslice)
			}
		}
		currSlices = nextSlices
	}
	return all
}

func diceSlices() [][]int {
	return [][]int{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}, []int{6}}
}

func rangeNums(start, end int) []int {
	var newSlice []int
	for idx := start; idx <= end; idx++ {
		newSlice = append(newSlice, idx)
	}
	return newSlice
}

func filterByLengthSum(targetSum, desiredLength int, permutations [][]int) [][]int {
	var selected [][]int
	for _, subslice := range permutations {
		if total(subslice) == targetSum && len(subslice) == desiredLength {
			selected = append(selected, subslice)
		}
	}
	return selected
}

func total(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
