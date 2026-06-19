package main

// LeetCode #3460: Longest Common Prefix After at Most One Removal
// https://leetcode.com/problems/longest-common-prefix-after-at-most-one-removal/
// Difficulty: Medium [Paid]
// Time: O(min(n,m)) Space: O(1)

import "fmt"

func longestCommonPrefix(s string, t string) int {
	n, m := len(s), len(t)
	maxLen := 0

	// without removal
	i, j := 0, 0
	for i < n && j < m && s[i] == t[j] {
		i++
		j++
	}
	maxLen = i

	// with one removal from s
	i, j = 0, 0
	removed := false
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else if !removed {
			i++
			removed = true
		} else {
			break
		}
	}
	if j > maxLen {
		maxLen = j
	}

	return maxLen
}

func main() {
	fmt.Println(longestCommonPrefix("abcde", "abfde")) // 2 (ab)
	fmt.Println(longestCommonPrefix("abc", "abc"))     // 3
	fmt.Println(longestCommonPrefix("abcd", "abxd"))   // 3 (abx vs abc — remove c from s -> abxd vs abxd... actually ab)
}
