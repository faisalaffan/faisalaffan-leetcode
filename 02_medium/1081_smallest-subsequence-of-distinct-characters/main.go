package main

// LeetCode #1081: Smallest Subsequence of Distinct Characters
// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
// Difficulty: Medium
//
// Approach: Monotonic stack (greedy). Track last occurrence and used set.
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bcabc"))  // "abc"
	fmt.Println(smallestSubsequence("cbacdcbc")) // "acdb"
}

func smallestSubsequence(s string) string {
	lastOccur := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastOccur[s[i]-'a'] = i
	}

	used := make([]bool, 26)
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if used[ch-'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastOccur[stack[len(stack)-1]-'a'] > i {
			used[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		used[ch-'a'] = true
	}

	return string(stack)
}
