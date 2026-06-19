package main

// LeetCode #96: Unique Binary Search Trees
// https://leetcode.com/problems/unique-binary-search-trees/
// Difficulty: Medium

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numTrees(3)) // 5

	// Test case 2
	fmt.Println(numTrees(1)) // 1

	// Test case 3
	fmt.Println(numTrees(4)) // 14
}

// Time: O(n^2) | Space: O(n)
