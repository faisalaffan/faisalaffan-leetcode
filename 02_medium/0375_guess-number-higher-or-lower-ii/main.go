package main

// LeetCode #375: Guess Number Higher or Lower II
// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
// Difficulty: Medium
// Time: O(n^3) | Space: O(n^2)

import "fmt"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for length := 2; length <= n; length++ {
		for start := 1; start <= n-length+1; start++ {
			end := start + length - 1
			dp[start][end] = 1<<31 - 1
			for pivot := start; pivot <= end; pivot++ {
				// Cost if pivot is wrong: pivot + max(cost of left, cost of right)
				cost := pivot + max(dp[start][pivot-1], dp[pivot+1][end])
				if cost < dp[start][end] {
					dp[start][end] = cost
				}
			}
		}
	}
	return dp[1][n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getMoneyAmount(10))
	// Expected: 16

	// Test case 2
	fmt.Println("Test 2:", getMoneyAmount(1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", getMoneyAmount(2))
	// Expected: 1
}
