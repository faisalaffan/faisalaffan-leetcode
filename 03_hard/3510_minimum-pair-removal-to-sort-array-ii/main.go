package main

// LeetCode #3510: Minimum Pair Removal to Sort Array II
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-ii/
// Difficulty: Hard
//
// Given an array, in each operation you can remove a pair of adjacent elements
// and insert their sum. Find the minimum number of operations to make the
// array non-decreasing.
//
// Approach: Greedy DP. Process from left to right, tracking the last value
// and the minimum cost to reach each position.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumPairRemoval([]int{1, 2, 3, 4, 5}))
	// Example 2
	fmt.Println(minimumPairRemoval([]int{5, 4, 3, 2, 1}))
	// Example 3
	fmt.Println(minimumPairRemoval([]int{1, 3, 2, 4}))
	// Edge: already sorted
	fmt.Println(minimumPairRemoval([]int{1, 2}))
	// Edge: single element
	fmt.Println(minimumPairRemoval([]int{1}))
}

func minimumPairRemoval(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// DP: dp[i] = min removals for prefix ending at i (inclusive)
	// lastVal[i] = value of last element after operations
	dp := make([]int, n)
	lastVal := make([]int, n)

	dp[0] = 0
	lastVal[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = 1 << 30
		// Try to take nums[i] as its own element
		for j := 0; j < i; j++ {
			// Sum of elements j+1..i as a single element after mergers
			sum := 0
			for k := j + 1; k <= i; k++ {
				sum += nums[k]
			}
			if lastVal[j]+dp[j] <= sum+dp[j]+int64(i-j-1) && lastVal[j] <= sum {
				ops := dp[j] + (i - j - 1)
				if ops < dp[i] {
					dp[i] = ops
					lastVal[i] = sum
				}
			}
		}
		// Also try merging nums[i] with adjacent to be >= lastVal
	}

	if dp[n-1] >= 1<<30 {
		return 0
	}
	return dp[n-1]
}
