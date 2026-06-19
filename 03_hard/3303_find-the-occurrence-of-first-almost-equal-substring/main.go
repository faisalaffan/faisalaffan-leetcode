package main

// LeetCode #3303: Find the Occurrence of First Almost Equal Substring
// https://leetcode.com/problems/find-the-occurrence-of-first-almost-equal-substring/
// Difficulty: Hard
//
// Given a string s and a pattern p, find the first index i in s such that
// the substring s[i:i+len(p)] is "almost equal" to p.
// A substring is "almost equal" to p if they can be made equal by
// modifying at most one character in either string.
// Return -1 if no such index exists.
//
// This is equivalent to finding the first position where the Hamming
// distance between s[i:i+m] and p is at most 1.
//
// For efficiency with large strings, we use Z-algorithm to compute
// longest common prefix and suffix matches, allowing O(n+m) time.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(firstAlmostEqualSubstring("abc", "abd"))
	// Example 2
	fmt.Println(firstAlmostEqualSubstring("hello", "world"))
	// Example 3: exact match
	fmt.Println(firstAlmostEqualSubstring("leetcode", "leet"))
	// Example 4: pattern longer than string
	fmt.Println(firstAlmostEqualSubstring("abc", "abcd"))
	// Example 5
	fmt.Println(firstAlmostEqualSubstring("abcdefgh", "abxdefgh"))
	// Example 6: single character
	fmt.Println(firstAlmostEqualSubstring("a", "b"))
}

func firstAlmostEqualSubstring(s string, p string) int {
	n, m := len(s), len(p)
	if m > n {
		return -1
	}
	if m == 0 {
		return 0
	}

	// Z-algorithm for longest common prefix.
	// combined = p + '#' + s gives us LCP of s[i:] with p at position m+1+i.
	combined := p + "#" + s
	z := zAlgo(combined)

	// LCP[i] = longest common prefix of s[i:] and p.
	lcp := make([]int, n)
	for i := 0; i < n; i++ {
		lcp[i] = z[m+1+i]
	}

	// Reverse for suffix matching.
	revP := reverse(p)
	revS := reverse(s)
	revCombined := revP + "#" + revS
	zRev := zAlgo(revCombined)

	// LCS[i] = longest common suffix of s[:i+m] and p.
	// For position i, the suffix starts at i+m-1 in s, which corresponds
	// to position (n-1)-(i+m-1) = n-i-m in the reversed string.
	lcs := make([]int, n)
	for i := 0; i <= n-m; i++ {
		revIdx := n - i - m // position in revS
		if revIdx >= 0 && revIdx < len(revS) {
			lcs[i] = zRev[m+1+revIdx]
		}
	}

	// Check each position.
	for i := 0; i <= n-m; i++ {
		// Characters matched from the start: lcp[i]
		// Characters matched from the end: lcs[i]
		// Gap = m - lcp[i] - lcs[i]
		// If gap <= 1, at most 1 character differs.
		if lcp[i]+lcs[i] >= m-1 {
			return i
		}
	}

	return -1
}

func zAlgo(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	return z
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Brute-force approach (for verification with small constraints).
func firstAlmostEqualSubstringBrute(s string, p string) int {
	n, m := len(s), len(p)
	if m > n {
		return -1
	}
	for i := 0; i <= n-m; i++ {
		diff := 0
		for j := 0; j < m; j++ {
			if s[i+j] != p[j] {
				diff++
				if diff > 1 {
					break
				}
			}
		}
		if diff <= 1 {
			return i
		}
	}
	return -1
}
