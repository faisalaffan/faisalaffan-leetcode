package main

// LeetCode #1342: Number of Steps to Reduce a Number to Zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// Difficulty: Easy
//
// LeetCode submission: func numberOfSteps(num int) int

import "fmt"

func main() {
	fmt.Println(NumberOfStepsToReduceANumberToZero(14)) // 6
	fmt.Println(NumberOfStepsToReduceANumberToZero(8))  // 4
	fmt.Println(NumberOfStepsToReduceANumberToZero(0))  // 0
}

// Time: O(log n), Space: O(1)
func NumberOfStepsToReduceANumberToZero(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
