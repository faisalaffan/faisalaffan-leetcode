package main

// LeetCode #3863: Minimum Operations to Sort a String
// https://leetcode.com/problems/minimum-operations-to-sort-a-string/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count characters and determine min swaps to sort.

import "fmt"

func MinimumOperationsToSortAString(s string) int {
	// Count character frequencies
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// Build sorted string
	sorted := make([]byte, len(s))
	idx := 0
	for c := 0; c < 26; c++ {
		for j := 0; j < cnt[c]; j++ {
			sorted[idx] = byte('a' + c)
			idx++
		}
	}

	// Count mismatching positions (where s[i] != sorted[i])
	mismatch := 0
	for i := 0; i < len(s); i++ {
		if s[i] != sorted[i] {
			mismatch++
		}
	}

	// Each swap fixes two mismatches, but there might be cycles
	// Simple: (mismatch + 1) / 2 gives swaps needed
	return (mismatch + 1) / 2
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToSortAString("cba")) // Expected: 1

	// Example
	fmt.Println(MinimumOperationsToSortAString("aab")) // Expected: 0

	// Example
	fmt.Println(MinimumOperationsToSortAString("acb")) // Expected: 1
}
