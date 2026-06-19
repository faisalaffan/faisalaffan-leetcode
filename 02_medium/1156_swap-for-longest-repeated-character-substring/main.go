package main

import (
	"fmt"
)

// LeetCode #1156: Swap For Longest Repeated Character Substring
// https://leetcode.com/problems/swap-for-longest-repeated-character-substring/
// Difficulty: Medium

// For each character, find the longest group. We can swap one character
// from elsewhere. Consider groups separated by exactly 1 char.

// Time: O(n)
// Space: O(n)

func maxRepOpt1(text string) int {
	n := len(text)

	// Count total frequency per character
	freq := make(map[byte]int)
	for i := 0; i < n; i++ {
		freq[text[i]]++
	}

	// Compute groups: (char, length)
	type group struct {
		char byte
		len  int
	}
	groups := make([]group, 0)

	i := 0
	for i < n {
		j := i
		for j < n && text[j] == text[i] {
			j++
		}
		groups = append(groups, group{text[i], j - i})
		i = j
	}

	maxLen := 0

	for idx, g := range groups {
		// Case 1: extend group by 1 if there's a spare of this char
		if freq[g.char] > g.len {
			if g.len+1 > maxLen {
				maxLen = g.len + 1
			}
		} else {
			if g.len > maxLen {
				maxLen = g.len
			}
		}

		// Case 2: merge two groups separated by exactly 1 char
		if idx > 0 && idx < len(groups)-1 &&
			groups[idx-1].char == groups[idx+1].char &&
			g.len == 1 {
			total := groups[idx-1].len + groups[idx+1].len
			if freq[groups[idx-1].char] > total {
				total++
			}
			if total > maxLen {
				maxLen = total
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 3)\n", maxRepOpt1("ababa"))
	fmt.Printf("%d (expected: 6)\n", maxRepOpt1("aaabaaa"))
	fmt.Printf("%d (expected: 4)\n", maxRepOpt1("aaaaa"))
}
