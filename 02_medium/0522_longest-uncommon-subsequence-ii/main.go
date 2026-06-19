package main

// LeetCode #522: Longest Uncommon Subsequence II
// https://leetcode.com/problems/longest-uncommon-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2 * L) where L is max length
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(FindLUSlength([]string{"aaa", "aaa", "aa"}))
}

func FindLUSlength(strs []string) int {
	maxLen := -1

	for i := 0; i < len(strs); i++ {
		isUnique := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSubseq(strs[i], strs[j]) {
				isUnique = false
				break
			}
		}
		if isUnique && len(strs[i]) > maxLen {
			maxLen = len(strs[i])
		}
	}

	return maxLen
}

func isSubseq(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
