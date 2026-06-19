package main

// LeetCode #2738: Count Occurrences in Text
// https://leetcode.com/problems/count-occurrences-in-text/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func CountOccurrencesInText(text string, word string) int {
	count := 0
	words := strings.Fields(text)
	for _, w := range words {
		if w == word {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountOccurrencesInText("hello world hello", "hello"))
	fmt.Println(CountOccurrencesInText("this is a test test this", "test"))
	fmt.Println(CountOccurrencesInText("unique", "none"))
}
