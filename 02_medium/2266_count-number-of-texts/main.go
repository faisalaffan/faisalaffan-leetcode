package main

// LeetCode #2266: Count Number of Texts
// https://leetcode.com/problems/count-number-of-texts/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] // press once
		// Check for multiple presses of same digit
		maxPress := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			maxPress = 4
		}
		for j := 2; j <= maxPress && j <= i; j++ {
			if pressedKeys[i-j] == pressedKeys[i-1] {
				dp[i] = (dp[i] + dp[i-j]) % mod
			} else {
				break
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(countTexts("22233"))
	// Expected: 8

	// Test case 2
	fmt.Println(countTexts("222222222222222222222222222222222222"))
	// Expected: 82876089

	// Test case 3
	fmt.Println(countTexts("33"))
	// Expected: 2
}
