package main

// LeetCode #1967: Number of Strings That Appear as Substrings in Word
// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "abc", "bc", "d"}, "abc")) // 3
	fmt.Println(NumberOfStringsThatAppearAsSubstringsInWord([]string{"a", "b", "c"}, "aaaaabbbbb"))  // 2
}

// Time: O(n * m), Space: O(1)
func NumberOfStringsThatAppearAsSubstringsInWord(patterns []string, word string) int {
	count := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
