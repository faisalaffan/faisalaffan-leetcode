package main

// LeetCode #3138: Minimum Length of Anagram Concatenation
// https://leetcode.com/problems/minimum-length-of-anagram-concatenation/
// Difficulty: Medium
// Time: O(n * sqrt(n)) | Space: O(n)

import "fmt"

func minAnagramLength(s string) int {
	n := len(s)

	freq := func(lo, hi int) [26]int {
		var f [26]int
		for i := lo; i < hi; i++ {
			f[s[i]-'a']++
		}
		return f
	}

	for l := 1; l <= n; l++ {
		if n%l != 0 {
			continue
		}
		base := freq(0, l)
		match := true
		for j := l; j < n && match; j += l {
			cur := freq(j, j+l)
			if cur != base {
				match = false
			}
		}
		if match {
			return l
		}
	}
	return n
}

func main() {
	fmt.Println(minAnagramLength("abba"))             // Expected: 2
	fmt.Println(minAnagramLength("abcabc"))           // Expected: 3
	fmt.Println(minAnagramLength("cdef"))             // Expected: 4
}
