package main

// LeetCode #316: Remove Duplicate Letters
// https://leetcode.com/problems/remove-duplicate-letters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func removeDuplicateLetters(s string) string {
	// Count last occurrence of each character
	lastOccur := [26]int{}
	for i := range s {
		lastOccur[s[i]-'a'] = i
	}

	stack := make([]byte, 0, len(s))
	seen := [26]bool{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if seen[ch-'a'] {
			continue
		}

		// Pop while stack top is greater and appears later
		for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]-'a'] {
			seen[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		seen[ch-'a'] = true
	}

	return string(stack)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeDuplicateLetters("bcabc"))
	// Expected: "abc"

	// Test case 2
	fmt.Println("Test 2:", removeDuplicateLetters("cbacdcbc"))
	// Expected: "acdb"

	// Test case 3: Single character
	fmt.Println("Test 3:", removeDuplicateLetters("aaaa"))
	// Expected: "a"
}
