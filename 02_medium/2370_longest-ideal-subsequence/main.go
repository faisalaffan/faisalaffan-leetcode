package main

// LeetCode #2370: Longest Ideal Subsequence
// https://leetcode.com/problems/longest-ideal-subsequence/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(26)
// DP over alphabet: dp[c] = longest ideal subsequence ending with char c.

import "fmt"

func main() {
	fmt.Println(longestIdealString("acfgbd", 2)) // 4
	fmt.Println(longestIdealString("abcd", 3))   // 4
}

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)
	var ans int
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		best := 0
		for p := 0; p < 26; p++ {
			if abs(c-p) <= k && dp[p] > best {
				best = dp[p]
			}
		}
		dp[c] = best + 1
		if dp[c] > ans {
			ans = dp[c]
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
