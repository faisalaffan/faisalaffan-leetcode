package main

import (
	"fmt"
)

// LeetCode #1190: Reverse Substrings Between Each Pair of Parentheses
// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
// Difficulty: Medium

// Find matching parentheses, then process from outside-in.
// Use wormhole approach: when hitting '(', jump to matching ')' and reverse direction.

// Time: O(n)
// Space: O(n)

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, 0)

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[open] = i
			pair[i] = open
		}
	}

	result := make([]byte, 0, n)
	dir := 1
	for i := 0; i < n; i += dir {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			dir = -dir
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(abcd)"), "dcba")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(u(love)i)"), "iloveu")
	fmt.Printf("%q (expected: %q)\n", reverseParentheses("(ed(et(oc))el)"), "leetcode")
}
