package main

// LeetCode #97: Interleaving String
// https://leetcode.com/problems/interleaving-string/
// Difficulty: Medium

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true

	// Test case 2
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false

	// Test case 3
	fmt.Println(isInterleave("", "", "")) // true
}

// Time: O(m*n) | Space: O(m*n)
