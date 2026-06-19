package main

// LeetCode #3407: Substring Matching Pattern
// https://leetcode.com/problems/substring-matching-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SubstringMatchingPattern("leetcode", "ee*e"))
	fmt.Println(SubstringMatchingPattern("car", "c*r"))
	fmt.Println(SubstringMatchingPattern("test", "t*t"))
}

// SubstringMatchingPattern checks if s matches pattern p where '*' matches any sequence of characters.
// Time: O(n * m). Space: O(n).
func SubstringMatchingPattern(s string, p string) bool {
	starIdx := -1
	for i, ch := range p {
		if ch == '*' {
			starIdx = i
			break
		}
	}

	left := p[:starIdx]
	right := p[starIdx+1:]

	// Left part must be prefix of some substring, right part must be suffix
	return strings.Contains(s, left) && strings.Contains(s, right) &&
		strings.Index(s, left) <= len(s)-len(right) &&
		strings.Index(s, left)+len(left) <= strings.LastIndex(s, right)
}
