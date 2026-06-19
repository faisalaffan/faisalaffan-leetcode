package main

// LeetCode #1682: Longest Palindromic Subsequence II
// https://leetcode.com/problems/longest-palindromic-subsequence-ii/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(n^2)

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i][j] = longest LPS length in s[i..j] with no equal adjacent chars
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				// Check adjacent condition: s[i] != the char that comes before/after
				if length == 2 {
					dp[i][j] = 2
				} else if s[i] != s[i+1] && s[j] != s[j-1] {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
				} else {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1])
				}
			}
			dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1]))
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbabab"))    // Expected: 4 ("baba" or "abab")
	fmt.Println(longestPalindromeSubseq("dcbccacdb")) // Expected: 4
	fmt.Println(longestPalindromeSubseq("a"))         // Expected: 1
}
