package main

// LeetCode #53: Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
// Difficulty: Medium

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum+nums[i] > nums[i] {
			currSum = currSum + nums[i]
		} else {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	// Test case 1
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	// Test case 2
	fmt.Println(maxSubArray([]int{1})) // 1

	// Test case 3
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8})) // 23
}

// Time: O(n) | Space: O(1)
