package main

// LeetCode #10: Regular Expression Matching
// https://leetcode.com/problems/regular-expression-matching/
// Difficulty: Hard

import "fmt"

// isMatch checks if string s matches pattern p.
// '.' matches any single character.
// '*' matches zero or more of the preceding element.
//
// Complexity: O(m*n) time, O(m*n) space where m = len(s), n = len(p)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true

	// handle patterns like a*, a*b*, a*b*c* matching empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// zero occurrences: skip the pattern char and *
				dp[i][j] = dp[i][j-2]
				// one or more occurrences: if preceding char matches s[i-1]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test: aa, a ->", isMatch("aa", "a"))           // false
	fmt.Println("Test: aa, a* ->", isMatch("aa", "a*"))          // true
	fmt.Println("Test: ab, .* ->", isMatch("ab", ".*"))          // true
	fmt.Println("Test: aab, c*a*b ->", isMatch("aab", "c*a*b")) // true
	fmt.Println("Test: mississippi, mis*is*p*. ->", isMatch("mississippi", "mis*is*p*.")) // false
}
