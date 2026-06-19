package main

// LeetCode #527: Word Abbreviation
// https://leetcode.com/problems/word-abbreviation/
// Difficulty: Hard

import (
	"fmt"
	"strconv"
)

func main() {
	input := []string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"}
	output := wordsAbbreviation(input)
	fmt.Println(output)
	// Expected: ["l2e","god","internal","me","i6t","interval","inte4n","f4e","intr4n"]
}

func wordsAbbreviation(words []string) []string {
	n := len(words)
	ans := make([]string, n)
	prefix := make([]int, n)

	for i := 0; i < n; i++ {
		prefix[i] = 1
		ans[i] = abbreviate(words[i], 1)
	}

	for i := 0; i < n; i++ {
		for {
			conflict := false
			for j := i + 1; j < n; j++ {
				if ans[i] == ans[j] {
					prefix[j]++
					ans[j] = abbreviate(words[j], prefix[j])
					conflict = true
				}
			}
			if conflict {
				prefix[i]++
				ans[i] = abbreviate(words[i], prefix[i])
			} else {
				break
			}
		}
	}
	return ans
}

func abbreviate(s string, k int) string {
	if k >= len(s)-2 {
		return s
	}
	abbr := s[:k] + strconv.Itoa(len(s)-k-1) + s[len(s)-1:]
	if len(abbr) >= len(s) {
		return s
	}
	return abbr
}
