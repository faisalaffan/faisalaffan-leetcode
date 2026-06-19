package main

// LeetCode #473: Matchsticks to Square
// https://leetcode.com/problems/matchsticks-to-square/
// Difficulty: Medium
// Time: O(4^n) worst case, but pruning makes it much faster
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MatchsticksToSquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(MatchsticksToSquare([]int{3, 3, 3, 3, 4}))
}

func MatchsticksToSquare(matchsticks []int) bool {
	sum := 0
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}
	target := sum / 4

	// Sort descending for better pruning
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	sides := make([]int, 4)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return sides[0] == target && sides[1] == target && sides[2] == target
		}
		for i := 0; i < 4; i++ {
			if sides[i]+matchsticks[idx] > target {
				continue
			}
			// Optimization: skip duplicate side lengths
			if i > 0 && sides[i] == sides[i-1] {
				continue
			}
			sides[i] += matchsticks[idx]
			if dfs(idx + 1) {
				return true
			}
			sides[i] -= matchsticks[idx]
		}
		return false
	}

	return dfs(0)
}
