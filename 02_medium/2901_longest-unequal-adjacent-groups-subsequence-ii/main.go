package main

// LeetCode #2901: Longest Unequal Adjacent Groups Subsequence II
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func LongestUnequalAdjacentGroupsSubsequenceIi(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	hamming := func(a, b string) int {
		if len(a) != len(b) {
			return -1
		}
		diff := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff
	}

	bestLen := 0
	bestIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if groups[j] != groups[i] && hamming(words[j], words[i]) == 1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
		if dp[i] > bestLen {
			bestLen = dp[i]
			bestIdx = i
		}
	}

	result := make([]string, bestLen)
	for i := bestLen - 1; i >= 0; i-- {
		result[i] = words[bestIdx]
		bestIdx = prev[bestIdx]
	}

	return result
}

func main() {
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceIi([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}))
}
