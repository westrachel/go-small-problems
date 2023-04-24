package main

import (
	"fmt"
)

// Write a function that accepts an integer "n" and returns the number of open lockers
// Details:
// > there are "n" number of lockers & "n" students
// > first student opens every locker
// > second student closes every second locker
// > third student toggles every third locker (if it's open they close it & vice versa)
// > fourth student toggles every fourth locker
// > function should return how many lockers are open after the "n" student completes
//    toggling every "n" locker
// > given input n is always >= 0

/* Examples:
n = 5
o o o o o 	// after student 1
o c o c o   // after student 2
o c c c o   // after student 3
o c c o o   // after student 4
o c c o c   // after student 5

n = 10
o o o o o o o o o o 	// after student 1
o c o c o c o c o c   // after student 2
o c c c o o o c c c   // after student 3
o c c o o o o o c c   // after student 4
o c c o c o o o c o   // after student 5
o c c o c c o o c o   // after student 6
o c c o c c c o c o   // after student 7
o c c o c c c c c o   // after student 8
o c c o c c c c o o   // after student 9
o c c o c c c c o c   // after student 10
total open: 3
*/

// Algorithm:
// i. initialize a list of size n that will contain all 'o''s to represent
//   n number of open lockers after the first student opens all the lockers
// ii. initialize a countOpen variable to keep track of the number of open
//      lockers - it should be initialized to "n" to represent the number of
//      open lockers after the first student walks through & opens all lockers
// iii. start a loop that keeps track of student counter that starts at 1 and
//     iterates to n - 1 and on each iteration:
//      > start an inner loop that has an index initialized to the current
//         student counter number and that iterates up to n -1 where the index
//         is incremented by the current student counter number plus 1
//            > the first iteration represents the "2nd" student but corresponds
//             with a student counter of 1 b/c arrays/slices indices start at 0
//              <=> this is why we need to increment the index by the student
//                counter plus 1
//          > on each inner loop iteration, toggle the status of the locker
//            at each index position visited in the slice of lockers
//              > if the locker at the index position is currently open 'o',
//                  then reassign it to 'c' and subtract 1 from countOpen
//              > if hte locker at the index position is currently closed 'c',
//                  then reassign it to 'o' and add 1 to countOpen
// iv. after iterating return countOpen

func main() {
	// Test cases:
	fmt.Println(openLockers(0))  // 0
	fmt.Println(openLockers(1))  // 1
	fmt.Println(openLockers(2))  // 1
	fmt.Println(openLockers(10)) // 3
	fmt.Println(openLockers(5))  // 2
}

func openLockers(n int) int {
	countOpen := n
	if n == 0 || n == 1 {
		return countOpen
	}
	lockers := initialLockers(n)

	for studentNum := 1; studentNum <= n-1; studentNum++ {
		for lockerIdx := studentNum; lockerIdx <= n-1; lockerIdx += (studentNum + 1) {
			if lockers[lockerIdx] == "o" {
				lockers[lockerIdx] = "c"
				countOpen--
			} else {
				lockers[lockerIdx] = "o"
				countOpen++
			}
		}
	}

	return countOpen
}

func initialLockers(n int) []string {
	var list []string
	for i := 0; i <= n-1; i++ {
		list = append(list, string('o'))
	}
	return list
}
