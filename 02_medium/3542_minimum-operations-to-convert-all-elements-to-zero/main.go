package main

// LeetCode #3542: Minimum Operations to Convert All Elements to Zero
// https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumOperationsToConvertAllElementsToZero([]int{1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", MinimumOperationsToConvertAllElementsToZero([]int{0, 0, 0}))
	// Test case 3
	fmt.Println("Test 3:", MinimumOperationsToConvertAllElementsToZero([]int{1, 1, 1}))
}

func MinimumOperationsToConvertAllElementsToZero(nums []int) int {
	ops := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			// flip from i onward (or flip just this element)
			ops++
		}
	}
	return ops
}
