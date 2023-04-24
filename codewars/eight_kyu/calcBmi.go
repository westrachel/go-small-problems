package main

import (
	"fmt"
	"math"
)

// Write a function that returns BMI info
// Reqs:
// > inputs: integer weight and height
//   > no input unit conversions needed: bmi = weight / height2
// > categories:
//   > bmi <= 18.5 return "Underweight"
//   > bmi <= 25.0 return "Normal"
//   > bmi <= 30.0 return "Overweight"
//   > bmi > 30 return "Obese"

func main() {
	// Test Cases:
	fmt.Println(calcBmi(1.54, 60))
	fmt.Println(calcBmi(1.65, 68))
	fmt.Println(calcBmi(1.3, 54))
}

func calcBmi(height float64, weight int) string {
	bmi := float64(weight) / math.Pow(height, 2)
	bmiStr := fmt.Sprintf("BMI: %.1f ", bmi)
	if bmi <= 18.5 {
		bmiStr += "Underweight"
	} else if bmi <= 25.0 {
		bmiStr += "Normal"
	} else if bmi <= 30.0 {
		bmiStr += "Overweight"
	} else {
		bmiStr += "Obese"
	}
	return bmiStr
}
