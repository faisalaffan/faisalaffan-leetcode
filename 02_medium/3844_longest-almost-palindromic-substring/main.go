package main

// LeetCode #3844: Longest Almost-Palindromic Substring
// https://leetcode.com/problems/longest-almost-palindromic-substring/
// Difficulty: Medium
// Time: O(N^2) | Space: O(1)
// Approach: Expand from each center. Track longest palindrome with at most
// one character removal (either via skipping a mismatch or by extending past
// the palindrome boundary by one).

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestAlmostPalindromicSubstring(s string) int {
	n := len(s)
	ans := 1

	// expand from center (l, r), allowing one skip
	check := func(l, r int) {
		// Phase 1: expand matching chars
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}

		pureLen := r - l - 1
		ans = max(ans, pureLen)

		// If we hit a boundary (l < 0 or r >= n), try including one more
		// char on the available side, then "removing" it.
		if l >= 0 && r >= n {
			// both sides have boundary: can't extend on either without breaking palindrome
			// but we can include the char at l and remove it
			ans = max(ans, pureLen+1)
		} else if l >= 0 && r < n {
			// left has char, right has char but they don't match (mismatch)
			// Try skipping left OR right, then continue expanding
			// skip left
			nl, nr := l-1, r
			for nl >= 0 && nr < n && s[nl] == s[nr] {
				nl--
				nr++
			}
			ans = max(ans, nr-nl-1)
			// skip right
			nl, nr = l, r+1
			for nl >= 0 && nr < n && s[nl] == s[nr] {
				nl--
				nr++
			}
			ans = max(ans, nr-nl-1)
		} else if l < 0 && r < n {
			// left boundary, right has char
			// include r and "remove" it
			ans = max(ans, pureLen+1)
		}
		// l < 0 && r >= n: both boundaries, nothing to add
	}

	for i := 0; i < n; i++ {
		check(i, i)   // odd length center
		check(i, i+1) // even length center
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestAlmostPalindromicSubstring("abca")) // Expected: 4

	// Example 2
	fmt.Println(LongestAlmostPalindromicSubstring("abba")) // Expected: 4

	// Example 3
	fmt.Println(LongestAlmostPalindromicSubstring("zzabba")) // Expected: 5
}
