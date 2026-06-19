package main

// LeetCode #784: Letter Case Permutation
// https://leetcode.com/problems/letter-case-permutation/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	curr := make([]byte, len(s))

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			result = append(result, string(curr))
			return
		}

		curr[idx] = s[idx]
		if unicode.IsLetter(rune(s[idx])) {
			curr[idx] = byte(unicode.ToLower(rune(s[idx])))
			backtrack(idx + 1)
			curr[idx] = byte(unicode.ToUpper(rune(s[idx])))
			backtrack(idx + 1)
		} else {
			backtrack(idx + 1)
		}
	}

	backtrack(0)
	return result
}
