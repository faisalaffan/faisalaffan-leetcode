package main

// LeetCode #1771: Maximize Palindrome Length From Subsequences
// https://leetcode.com/problems/maximize-palindrome-length-from-subsequences/
// Difficulty: Hard

import "fmt"

func main() {
	word1 := "cacb"
	word2 := "cbba"
	fmt.Printf("Test 1 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 5)\n\n", longestPalindrome(word1, word2))

	word1 = "ab"
	word2 = "ab"
	fmt.Printf("Test 2 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 3)\n\n", longestPalindrome(word1, word2))

	word1 = "aa"
	word2 = "bb"
	fmt.Printf("Test 3 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 0)\n\n", longestPalindrome(word1, word2))

	word1 = "cebdedc"
	word2 = "d"
	fmt.Printf("Test 4 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 5)\n", longestPalindrome(word1, word2))
}

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	n := len(s)
	n1 := len(word1)

	// dp[i][j] = longest palindromic subsequence in s[i..j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	result := 0

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				// Must pick at least one char from each word
				if i < n1 && j >= n1 && dp[i][j] > result {
					result = dp[i][j]
				}
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return result
}
