package main

import (
	"fmt"
	"strings"
)

// Write a function that takes a string and alternates the case
// of each alphabetical character in the string.

func main() {
	// Test cases:
	fmt.Println(alternateCase("hello world") == "HELLO WORLD")
	fmt.Println(alternateCase("HELLO WORLD") == "hello world")
	fmt.Println(alternateCase("HeLLo WoRLD") == "hEllO wOrld")
	fmt.Println(alternateCase("hello WORLD") == "HELLO world")
	fmt.Println(alternateCase("12345") == "12345")
	fmt.Println(alternateCase("1a2b3c4d5e") == "1A2B3C4D5E")
}

func alternateCase(str string) string {
	var altCaseStr string
	for i := 0; i < len(str); i++ {
		var char = str[i : i+1]
		if strings.ToUpper(char) == char {
			altCaseStr += strings.ToLower(char)
		} else {
			altCaseStr += strings.ToUpper(str[i : i+1])
		}
	}
	return altCaseStr
}
