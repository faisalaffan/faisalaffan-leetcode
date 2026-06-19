package main

// LeetCode #2300: Successful Pairs of Spells and Potions
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
// Difficulty: Medium
// Time: O((n + m) log m) | Space: O(1)

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	result := make([]int, len(spells))

	for i, s := range spells {
		need := (int(success) + s - 1) / s // ceiling division
		idx := sort.SearchInts(potions, need)
		result[i] = m - idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	// Expected: [4, 0, 3]

	// Test case 2
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	// Expected: [2, 0, 2]
}
