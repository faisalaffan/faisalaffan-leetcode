package main

// LeetCode #921: Minimum Add to Make Parentheses Valid
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func minAddToMakeValid(s string) int {
	open, add := 0, 0
	for _, ch := range s {
		if ch == '(' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				add++
			}
		}
	}
	return add + open
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()"))
	fmt.Println(minAddToMakeValid("()))(("))
}
