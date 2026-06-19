package main

// LeetCode #3503: Longest Palindrome After Substring Concatenation I
// https://leetcode.com/problems/longest-palindrome-after-substring-concatenation-i/
// Difficulty: Medium
// Complexity: O(n^2 * m^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", LongestPalindromeAfterSubstringConcatenationI("ab", "ba"))
	// Test case 2
	fmt.Println("Test 2:", LongestPalindromeAfterSubstringConcatenationI("a", "a"))
	// Test case 3
	fmt.Println("Test 3:", LongestPalindromeAfterSubstringConcatenationI("abc", "cba"))
}

func LongestPalindromeAfterSubstringConcatenationI(s string, t string) int {
	isPal := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	maxLen := 0
	// Try all substrings of s concatenated with all substrings of t
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			for p := 0; p <= len(t); p++ {
				for q := p; q <= len(t); q++ {
					candidate := s[i:j] + t[p:q]
					if isPal(candidate) && len(candidate) > maxLen {
						maxLen = len(candidate)
					}
				}
			}
		}
	}
	return maxLen
}
