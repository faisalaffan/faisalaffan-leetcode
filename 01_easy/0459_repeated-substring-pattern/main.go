package main

// LeetCode #459: Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func RepeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	fmt.Println(RepeatedSubstringPattern("abab"))
	fmt.Println(RepeatedSubstringPattern("aba"))
	fmt.Println(RepeatedSubstringPattern("abcabcabcabc"))
}
