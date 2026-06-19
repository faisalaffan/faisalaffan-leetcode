package main

// LeetCode #3926: Count Valid Word Occurrences
// https://leetcode.com/problems/count-valid-word-occurrences/
// Difficulty: Medium
// Time: O(N + Q) | Space: O(N)
// Approach: Concatenate chunks, extract words (lowercase letters +
// joiner hyphens), count with hash map, look up queries.

import (
	"fmt"
	"strings"
)

func CountValidWordOccurrences(chunks []string, queries []string) []int {
	s := strings.Join(chunks, "")

	wordCount := make(map[string]int)
	n := len(s)
	i := 0
	for i < n {
		// Skip non-word chars (spaces, leading hyphens, isolated hyphens)
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '-' {
			i++
			continue
		}

		// Start of a word
		j := i
		for j < n {
			ch := s[j]
			if ch >= 'a' && ch <= 'z' {
				j++
				continue
			}
			if ch == '-' && j+1 < n && s[j+1] >= 'a' && s[j+1] <= 'z' {
				// Joiner hyphen: part of current word
				j += 2
				continue
			}
			break
		}

		if j > i {
			wordCount[s[i:j]]++
		}
		i = j
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		ans[idx] = wordCount[q]
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountValidWordOccurrences(
		[]string{"hello wor", "ld hello"},
		[]string{"hello", "world", "wor"},
	)) // Expected: [2 1 0]

	// Example 2
	fmt.Println(CountValidWordOccurrences(
		[]string{"a-b a--b ", "a-", "b"},
		[]string{"a-b", "a", "b"},
	)) // Expected: [2 1 1]

	// Example 3
	fmt.Println(CountValidWordOccurrences(
		[]string{"-cat dog- mouse"},
		[]string{"cat", "dog", "mouse", "cat-dog"},
	)) // Expected: [1 1 1 0]
}
