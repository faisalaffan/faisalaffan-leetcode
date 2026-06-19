package main

// LeetCode #2707: Extra Characters in a String
// https://leetcode.com/problems/extra-characters-in-a-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	wordSet := make(map[string]bool)
	for _, w := range dictionary {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		// Try to match word ending at i-1
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if wordSet[s[j:i]] {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	// Expected: 1 ('s' is extra)

	// Test case 2
	fmt.Println("Test 2:", minExtraChar("sayhelloworld", []string{"hello", "world"}))
	// Expected: 3 ('say' is extra)

	// Test case 3
	fmt.Println("Test 3:", minExtraChar("abcd", []string{"a", "b", "c", "d"}))
	// Expected: 0
}
