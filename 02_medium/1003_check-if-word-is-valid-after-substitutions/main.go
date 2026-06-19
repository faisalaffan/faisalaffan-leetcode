package main

// LeetCode #1003: Check If Word Is Valid After Substitutions
// https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
// Difficulty: Medium
//
// Approach: Use a stack. When we see "c", check if top two are "a" and "b".
//           If so, they form "abc" and we pop them. Otherwise, push "c".
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isValid("aabcbc"))  // true
	fmt.Println(isValid("abcabcababcc")) // true
	fmt.Println(isValid("abccba"))  // false
	fmt.Println(isValid("cababc"))  // false
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == 'c' {
			n := len(stack)
			if n >= 2 && stack[n-1] == 'b' && stack[n-2] == 'a' {
				stack = stack[:n-2]
				continue
			}
		}
		stack = append(stack, s[i])
	}

	return len(stack) == 0
}
