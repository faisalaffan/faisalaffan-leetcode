package main

// LeetCode #1035: Uncrossed Lines
// https://leetcode.com/problems/uncrossed-lines/
// Difficulty: Medium
//
// Approach: DP - Longest Common Subsequence (LCS)
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))    // 2
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})) // 3
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
