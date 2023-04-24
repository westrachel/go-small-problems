package main

import (
	"fmt"
)

func main() {
	// Test cases:
	fmt.Println(triangleNums(1)) // 1
	fmt.Println(triangleNums(2)) // 3
	fmt.Println(triangleNums(6)) // 21
}

func triangleNums(n int) int {
	sum := 0
	for naturalNum := 0; naturalNum <= n; naturalNum++ {
		sum += naturalNum
	}
	return sum
}
