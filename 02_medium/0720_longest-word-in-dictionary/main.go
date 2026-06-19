package main

// LeetCode #720: Longest Word in Dictionary
// https://leetcode.com/problems/longest-word-in-dictionary/
// Difficulty: Medium
// Time: O(n * L + n log n)
// Space: O(n * L)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}

func longestWord(words []string) string {
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}

	sort.Strings(words)

	result := ""
	for _, w := range words {
		if len(w) <= len(result) {
			continue
		}
		valid := true
		for i := 1; i < len(w); i++ {
			if !wordSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			result = w
		}
	}

	return result
}
