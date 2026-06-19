package main

// LeetCode #521: Longest Uncommon Subsequence I
// https://leetcode.com/problems/longest-uncommon-subsequence-i/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func LongestUncommonSubsequenceI(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(LongestUncommonSubsequenceI("aba", "cdc"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "bbb"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "aaa"))
}
