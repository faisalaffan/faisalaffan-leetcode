package main

// LeetCode #139: Word Break
// https://leetcode.com/problems/word-break/
// Difficulty: Medium

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	// Test case 1
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"})) // true

	// Test case 2
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"})) // true

	// Test case 3
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

// Time: O(n^2) | Space: O(n)
