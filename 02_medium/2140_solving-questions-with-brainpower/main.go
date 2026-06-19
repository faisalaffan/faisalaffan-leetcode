package main

// LeetCode #2140: Solving Questions With Brainpower
// https://leetcode.com/problems/solving-questions-with-brainpower/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		points := int64(questions[i][0])
		brainpower := questions[i][1]
		// Skip this question
		best := dp[i+1]
		// Take this question
		next := i + brainpower + 1
		take := points
		if next <= n {
			take += dp[next]
		}
		if take > best {
			best = take
		}
		dp[i] = best
	}

	return dp[0]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", mostPoints([][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}))
	// Expected: 157
}
