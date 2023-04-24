package main

import (
	"fmt"
)

/*
  Write a function that takes a tic-tac-toe board (3x3 array) and returns
  a number indicating whether the game has a winner or not

Ex:
 Input: 			Output:
 [[0, 0, 1],			> -1
 [0, 1, 2],
 [2, 1, 0]]

  > O represents empty space
  > 1 represents an X
  > 2 represents an O

 Possible outputs:
  > -1 if game isn't finished yet (there are empty spaces)
  > 0 if it's a tie
  > 1 if X has won
  > 2 if O has won

  Ways to Win:
  > integer positions in the board range from 0 -> 8
  > 3 in a row horizontally: so same character (1 or 2) at positions:
        rows: 0, 0, 0   columns: 0, 1, 2
        rows: 1, 1, 1   columns: 0, 1, 2
        rows: 2, 2, 2   columns: 0, 1, 2
  > 3 in a row vertically: so same character (1 or 2) at positions:
        rows: 0, 1, 2  columns: 0, 0, 0
        rows: 0, 1, 2  columns: 1, 1, 1
        rows: 0, 1, 1  columns: 2, 2, 2
  > 3 in a row diagonally: so same character (1 or 2) at integer positions: [2, 4, 6], [0, 4, 8]
        rows: 0, 1, 2  columns: 0, 1, 2
        rows: 0, 1, 2  columns: 2, 1, 0

  Algorithm:
  i. declare a diagonal winner
       > check if the following elements are equivalent:
         [0][0] == [1][1] == [2][2]
           > if they are, then return one of the elements as long as its not 0
       > check if [2][0] == [1][1] == [0][2] are equivalent and are nonzero
           > if this is true then return one of the elements
       > in following steps will checked if there are any empty spaces so return -100
          if neither of the prior conditions are true
  ii. declare a horizontal row winner by iterating over the rows and for each row:
         > declare a rowVal variable assigned to the 0th indexed element
         > if the elements at the 1st and 2nd index are equal to the element at the 0th
            index then return rowVal as long as it is not 0 (b/c 0 doesn't indicate a winner)
         >  if any of the characters are 0s return -1
         > otherwise return 0
  iii. declare a vertical row winner by iterating from colIdx 0 to 2 and on each iteration:
         > check if the elements at colIdx in rows 0, 1, and 2 are equivalent and are not 0
            > if they are return one of the elements, bc have a winner
         > if they aren't and any of them are 0 then return -1
         > otherwise return 0
  iv. return the max value of the horizontal, vertical, and diagonal row winners
*/

func main() {
	// Test cases:
	unfinished := [3][3]int{[3]int{0, 0, 1}, [3]int{0, 1, 2}, [3]int{2, 1, 0}}
	fmt.Println(IsSolved(unfinished)) // -1

	horizWinner := [3][3]int{[3]int{1, 1, 1}, [3]int{0, 1, 2}, [3]int{2, 0, 0}}
	fmt.Println(IsSolved(horizWinner)) // 1

	vertWinner := [3][3]int{[3]int{1, 1, 2}, [3]int{0, 0, 2}, [3]int{2, 1, 2}}
	fmt.Println(IsSolved(vertWinner)) // 2

	diagWinner := [3][3]int{[3]int{1, 1, 2}, [3]int{0, 1, 2}, [3]int{2, 2, 1}}
	fmt.Println(IsSolved(diagWinner)) // 1

	tie := [3][3]int{[3]int{2, 2, 1}, [3]int{1, 1, 2}, [3]int{2, 2, 1}}
	fmt.Println(IsSolved(tie)) // 0

	horizWinner = [3][3]int{[3]int{1, 1, 1}, [3]int{0, 2, 2}, [3]int{0, 0, 0}}
	fmt.Println(IsSolved(horizWinner)) // 1

	unfinished = [3][3]int{[3]int{1, 2, 2}, [3]int{2, 1, 1}, [3]int{0, 0, 0}}
	fmt.Println(IsSolved(unfinished)) // -1

	unfinished = [3][3]int{[3]int{0, 2, 1}, [3]int{2, 0, 2}, [3]int{1, 1, 0}}
	fmt.Println(IsSolved(unfinished)) // -1
}

func IsSolved(board [3][3]int) int {
	diagResult := diagonalWinner(board)
	horizVertResult := horizVertWinner(board)

	if horizVertResult > diagResult {
		return horizVertResult
	} else {
		return diagResult
	}
}

func diagonalWinner(board [3][3]int) int {
	middle := board[1][1]
	firstDiag := board[0][0] == middle && middle == board[2][2]
	secondDiag := board[0][2] == middle && middle == board[2][0]

	if (firstDiag || secondDiag) && middle != 0 {
		return middle
	} else {
		return -100
	}
}
func horizVertWinner(board [3][3]int) int {
	countZeros := 0
	for idx := 0; idx <= 2; idx++ {
		rows := findVertHorizRows(board, idx)

		for _, row := range rows {
			middle := row[1]
			if row[0] == middle && middle == row[2] && middle != 0 {
				return middle
			} else if includesZero(row) {
				countZeros += 1
			}
		}
	}
	return decideResult(countZeros)
}

func findVertHorizRows(board [3][3]int, idx int) [2][3]int {
	vertical := [3]int{board[0][idx], board[1][idx], board[2][idx]}
	horizontal := board[idx]
	return [2][3]int{vertical, horizontal}
}

func decideResult(countZeros int) int {
	if countZeros > 0 {
		return -1
	} else {
		return 0
	}
}

func includesZero(arr [3]int) bool {
	return arr[0] == 0 || arr[1] == 0 || arr[2] == 0
}
