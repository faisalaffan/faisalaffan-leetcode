package main

// LeetCode #686: Repeated String Match
// https://leetcode.com/problems/repeated-string-match/
// Difficulty: Medium
// Time: O(n * m) worst case
// Space: O(n + m)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("abc", "wxyz"))
}

func repeatedStringMatch(a string, b string) int {
	maxRepeats := len(b)/len(a) + 3
	var sb strings.Builder

	for i := 1; i <= maxRepeats; i++ {
		sb.WriteString(a)
		if strings.Contains(sb.String(), b) {
			return i
		}
	}

	return -1
}
