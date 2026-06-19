package main

// LeetCode #678: Valid Parenthesis String
// https://leetcode.com/problems/valid-parenthesis-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
	low, high := 0, 0

	for _, c := range s {
		if c == '(' {
			low++
			high++
		} else if c == ')' {
			low--
			high--
		} else { // '*'
			low--
			high++
		}

		if high < 0 {
			return false
		}
		if low < 0 {
			low = 0
		}
	}

	return low == 0
}
