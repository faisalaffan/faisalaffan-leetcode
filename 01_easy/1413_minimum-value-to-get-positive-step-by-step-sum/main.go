package main

// LeetCode #1413: Minimum Value to Get Positive Step by Step Sum
// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// Difficulty: Easy
//
// LeetCode submission: func minStartValue(nums []int) int

import "fmt"

func main() {
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{-3, 2, -3, 4, 2})) // 5
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, 2}))             // 1
	fmt.Println(MinimumValueToGetPositiveStepByStepSum([]int{1, -2, -3}))        // 5
}

// Time: O(n), Space: O(1)
func MinimumValueToGetPositiveStepByStepSum(nums []int) int {
	minSum, sum := 0, 0
	for _, v := range nums {
		sum += v
		if sum < minSum {
			minSum = sum
		}
	}
	return -minSum + 1
}
