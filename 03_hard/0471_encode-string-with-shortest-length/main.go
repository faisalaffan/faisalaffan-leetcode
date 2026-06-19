package main

// LeetCode #471: Encode String with Shortest Length
// https://leetcode.com/problems/encode-string-with-shortest-length/
// Difficulty: Hard [Paid]
// Approach: DP interval encoding. dp[i][j] = shortest encoded form of s[i:j+1].
// For each substring, try to compress it using KMP-like period detection,
// or split it into two parts and combine.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("471 - Encode String with Shortest Length")

	// Test cases
	fmt.Printf("encode(\"aaa\") = %q (expected: \"aaa\")\n", encode("aaa"))
	fmt.Printf("encode(\"aaaaa\") = %q (expected: \"5[a]\")\n", encode("aaaaa"))
	fmt.Printf("encode(\"aabcaabcd\") = %q (expected: \"2[aabc]d\")\n", encode("aabcaabcd"))
	fmt.Printf("encode(\"abbbabbbcabbbabbbc\") = %q (expected: \"2[2[abbb]c]\")\n", encode("abbbabbbcabbbabbbc"))
	fmt.Printf("encode(\"aaaaaaaaaa\") = %q (expected: \"10[a]\")\n", encode("aaaaaaaaaa"))
	fmt.Printf("encode(\"\") = %q (expected: \"\")\n", encode(""))
	fmt.Printf("encode(\"a\") = %q (expected: \"a\")\n", encode("a"))
	fmt.Printf("encode(\"ab\") = %q (expected: \"ab\")\n", encode("ab"))
	fmt.Printf("encode(\"abcabcabc\") = %q (expected: \"3[abc]\")\n", encode("abcabcabc"))
	fmt.Printf("encode(\"abbbabbbc\") = %q (expected: \"abbbabbbc\" (no compression -> \"a2[bbb]c\" vs original, original shorter))\n", encode("abbbabbbc"))
}

func encode(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// dp[i][j] = shortest encoded form of s[i:j+1]
	dp := make([][]string, n)
	for i := range dp {
		dp[i] = make([]string, n)
		dp[i][i] = string(s[i])
	}

	// length from 2 to n
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			substr := s[i : j+1]

			// Option 1: try to compress the whole substring
			// Find the smallest repeating pattern
			encoded := substr
			period := findPeriod(substr)
			if period > 0 && len(substr) > period {
				encoded = fmt.Sprintf("%d[%s]", len(substr)/period, dp[i][i+period-1])
			}

			// Option 2: split into two parts
			for k := i; k < j; k++ {
				combined := dp[i][k] + dp[k+1][j]
				if len(combined) < len(encoded) {
					encoded = combined
				}
			}

			// If the original is shorter, keep the original
			if len(substr) < len(encoded) {
				encoded = substr
			}

			dp[i][j] = encoded
		}
	}

	return dp[0][n-1]
}

// findPeriod returns the smallest period of the string.
// If the string can be formed by repeating a prefix, returns the length of that prefix.
// Otherwise returns 0.
func findPeriod(s string) int {
	n := len(s)
	// Try all possible periods from 1 to n/2
	for p := 1; p <= n/2; p++ {
		if n%p != 0 {
			continue
		}
		pattern := s[:p]
		if strings.Repeat(pattern, n/p) == s {
			return p
		}
	}
	return 0
}

