package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1181: Before and After Puzzle
// https://leetcode.com/problems/before-and-after-puzzle/
// Difficulty: Medium [Paid]

// Given list of phrases. Merge phrase i and j if last word of i
// equals first word of j. Result: i + j[firstWordLen:].
// Return sorted unique results.

// Time: O(n^2 * L)
// Space: O(n^2)

func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)
	firstWords := make([]string, n)
	lastWords := make([]string, n)
	words := make([][]string, n)

	for i, p := range phrases {
		words[i] = strings.Fields(p)
		firstWords[i] = words[i][0]
		lastWords[i] = words[i][len(words[i])-1]
	}

	seen := make(map[string]bool)
	result := make([]string, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if lastWords[i] == firstWords[j] {
				merged := phrases[i] + phrases[j][len(firstWords[j]):]
				if !seen[merged] {
					seen[merged] = true
					result = append(result, merged)
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [writing code rocks])\n",
		beforeAndAfterPuzzles([]string{"writing code", "code rocks"}))

	fmt.Printf("%v (expected: [a d a b c d])\n",
		beforeAndAfterPuzzles([]string{"a b", "b c", "c d"}))
}
