package main

// LeetCode #2268: Minimum Number of Keypresses
// https://leetcode.com/problems/minimum-number-of-keypresses/
// Difficulty: Medium [Paid]
// Time: O(n + 26 log 26) | Space: O(26)

import (
	"fmt"
	"sort"
)

func minimumKeypresses(s string) int {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	presses := 0
	for i, c := range count {
		if c == 0 {
			break
		}
		presses += c * (i/9 + 1)
	}
	return presses
}

func main() {
	// Test case 1
	fmt.Println(minimumKeypresses("apple"))
	// Expected: 5

	// Test case 2
	fmt.Println(minimumKeypresses("abcdefghijkl"))
	// Expected: 15

	// Test case 3
	fmt.Println(minimumKeypresses("aaaaaaa"))
	// Expected: 7
}
