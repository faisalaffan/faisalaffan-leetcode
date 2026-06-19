package main

// LeetCode #796: Rotate String
// https://leetcode.com/problems/rotate-string/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
}

// rotateString checks if goal can be obtained by rotating s.
// Time: O(n). Space: O(n).
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
