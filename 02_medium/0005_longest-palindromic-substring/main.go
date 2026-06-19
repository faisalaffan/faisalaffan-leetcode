package main

// LeetCode #5: Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/
// Difficulty: Medium

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		expandAroundCenter(i, i)   // odd length
		expandAroundCenter(i, i+1) // even length
	}

	return s[start : start+maxLen]
}

func main() {
	// Test case 1
	fmt.Println(longestPalindrome("babad")) // "bab" or "aba"

	// Test case 2
	fmt.Println(longestPalindrome("cbbd")) // "bb"

	// Test case 3
	fmt.Println(longestPalindrome("a")) // "a"
}

// Time: O(n^2) | Space: O(1)
