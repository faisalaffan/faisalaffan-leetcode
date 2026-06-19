package main

import (
	"fmt"
)

// LeetCode #1239: Maximum Length of a Concatenated String with Unique Characters
// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
// Difficulty: Medium

// Backtracking with bitmask. Each string is represented as a bitmask.
// Only concatenate if no character conflict.

// Time: O(2^n) worst case
// Space: O(n)

func maxLength(arr []string) int {
	masks := make([]int, 0)

	for _, s := range arr {
		mask := 0
		valid := true
		for _, ch := range s {
			bit := 1 << (ch - 'a')
			if mask&bit != 0 {
				valid = false
				break
			}
			mask |= bit
		}
		if valid {
			masks = append(masks, mask)
		}
	}

	maxLen := 0
	var backtrack func(idx, mask, length int)
	backtrack = func(idx, mask, length int) {
		if length > maxLen {
			maxLen = length
		}
		for i := idx; i < len(masks); i++ {
			if mask&masks[i] == 0 {
				backtrack(i+1, mask|masks[i], length+countBits(masks[i]))
			}
		}
	}

	backtrack(0, 0, 0)
	return maxLen
}

func countBits(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maxLength([]string{"un", "iq", "ue"}))
	fmt.Printf("%d (expected: 6)\n", maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Printf("%d (expected: 26)\n", maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
}
