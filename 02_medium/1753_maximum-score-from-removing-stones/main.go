package main

// LeetCode #1753: Maximum Score From Removing Stones
// https://leetcode.com/problems/maximum-score-from-removing-stones/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumScore(a int, b int, c int) int {
	piles := []int{a, b, c}
	sort.Ints(piles)
	// If the largest pile is >= sum of other two, we can only take sum of those two
	if piles[2] >= piles[0]+piles[1] {
		return piles[0] + piles[1]
	}
	// Otherwise, we can take (a+b+c)/2 stones
	return (a + b + c) / 2
}

func main() {
	fmt.Println(maximumScore(2, 4, 6)) // Expected: 6
	fmt.Println(maximumScore(4, 4, 6)) // Expected: 7
	fmt.Println(maximumScore(1, 8, 8)) // Expected: 8
}
