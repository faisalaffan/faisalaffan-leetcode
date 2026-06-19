package main

// LeetCode #3909: Compare Sums of Bitonic Parts
// https://leetcode.com/problems/compare-sums-of-bitonic-parts/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Find peak where nums[i] > nums[i+1]. Sum left part from 0 to peak,
// sum right part from peak to n-1. Return -1/0/1.

import "fmt"

func CompareSumsOfBitonicParts(nums []int) int {
	n := len(nums)
	peak := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			peak = i
			break
		}
	}

	leftSum := 0
	for i := 0; i <= peak; i++ {
		leftSum += nums[i]
	}
	rightSum := 0
	for i := peak; i < n; i++ {
		rightSum += nums[i]
	}

	if leftSum > rightSum {
		return 0
	} else if rightSum > leftSum {
		return 1
	}
	return -1
}

func main() {
	// Example 1
	fmt.Println(CompareSumsOfBitonicParts([]int{1, 3, 2, 1})) // Expected: 1

	// Example 2
	fmt.Println(CompareSumsOfBitonicParts([]int{2, 4, 5, 2})) // Expected: 0

	// Example 3
	fmt.Println(CompareSumsOfBitonicParts([]int{1, 2, 4, 3})) // Expected: -1
}
