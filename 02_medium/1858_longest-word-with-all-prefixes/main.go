package main

// LeetCode #1858: Longest Word With All Prefixes
// https://leetcode.com/problems/longest-word-with-all-prefixes/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestWord([]string{"k", "ki", "kir", "kira", "kiran"}))
	fmt.Println(LongestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(LongestWord([]string{"abc", "ab", "a"}))
}

// Time: O(n log n + total chars), Space: O(total unique prefixes)
func LongestWord(words []string) string {
	prefixSet := make(map[string]bool)
	for _, w := range words {
		prefixSet[w] = true
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) > len(words[j])
		}
		return words[i] < words[j]
	})

	for _, w := range words {
		valid := true
		for i := 1; i <= len(w); i++ {
			if !prefixSet[w[:i]] {
				valid = false
				break
			}
		}
		if valid {
			return w
		}
	}
	return ""
}
