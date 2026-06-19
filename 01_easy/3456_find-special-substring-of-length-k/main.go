package main

// LeetCode #3456: Find Special Substring of Length K
// https://leetcode.com/problems/find-special-substring-of-length-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindSpecialSubstringOfLengthK("aaabaaa", 3))
	fmt.Println(FindSpecialSubstringOfLengthK("abc", 2))
}

// FindSpecialSubstringOfLengthK returns true if there is a substring of length k consisting of a single character, surrounded by different characters (or boundaries).
// Time: O(n). Space: O(1).
func FindSpecialSubstringOfLengthK(s string, k int) bool {
	n := len(s)
	for i := 0; i <= n-k; i++ {
		same := true
		for j := i; j < i+k-1; j++ {
			if s[j] != s[j+1] {
				same = false
				break
			}
		}
		if !same {
			continue
		}
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		if i+k < n && s[i+k] == s[i] {
			continue
		}
		return true
	}
	return false
}
