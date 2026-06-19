package main

// LeetCode #2730: Find the Longest Semi-Repetitive Substring
// https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func FindTheLongestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	maxLen := 0
	left := 0
	lastPair := -1

	for right := 1; right < n; right++ {
		if s[right] == s[right-1] {
			if lastPair != -1 {
				left = lastPair
			}
			lastPair = right
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("52233"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("0001"))
	fmt.Println(FindTheLongestSemiRepetitiveSubstring("1111111"))
}
