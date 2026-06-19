package main

// LeetCode #377: Combination Sum IV
// https://leetcode.com/problems/combination-sum-iv/
// Difficulty: Medium
// Time: O(target * n) | Space: O(target)

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", combinationSum4([]int{1, 2, 3}, 4))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", combinationSum4([]int{9}, 3))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", combinationSum4([]int{1, 2, 3}, 3))
	// Expected: 4
}
