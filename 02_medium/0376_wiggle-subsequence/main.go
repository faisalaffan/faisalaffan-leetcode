package main

// LeetCode #376: Wiggle Subsequence
// https://leetcode.com/problems/wiggle-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// Expected: 2
}
