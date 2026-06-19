package main

import (
	"fmt"
)

// LeetCode #1234: Replace the Substring for Balanced String
// https://leetcode.com/problems/replace-the-substring-for-balanced-string/
// Difficulty: Medium

// Sliding window. Find smallest substring such that outside it,
// each of Q,W,E,R appears at most n/4 times.

// Time: O(n)
// Space: O(1)

func balancedString(s string) int {
	n := len(s)
	target := n / 4
	count := make(map[byte]int)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}

	// Check if already balanced
	balanced := true
	for _, c := range []byte{'Q', 'W', 'E', 'R'} {
		if count[c] > target {
			balanced = false
			break
		}
	}
	if balanced {
		return 0
	}

	left := 0
	minLen := n

	for right := 0; right < n; right++ {
		count[s[right]]--

		for left <= right {
			ok := true
			for _, c := range []byte{'Q', 'W', 'E', 'R'} {
				if count[c] > target {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			count[s[left]]++
			left++
		}
	}

	return minLen
}

func main() {
	fmt.Printf("%d (expected: 0)\n", balancedString("QWER"))
	fmt.Printf("%d (expected: 1)\n", balancedString("QQWE"))
	fmt.Printf("%d (expected: 2)\n", balancedString("QQQW"))
}
