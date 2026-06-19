package main

// LeetCode #2896: Apply Operations to Make Two Strings Equal
// https://leetcode.com/problems/apply-operations-to-make-two-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeTwoStringsEqual(s1 string, s2 string, x int) int {
	// Find positions where characters differ
	diff := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff)%2 != 0 {
		return -1
	}
	if len(diff) == 0 {
		return 0
	}

	m := len(diff)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}

		// Option 1: flip s[l] and s[l+1] (cost = diff[l+1] - diff[l])
		best := solve(l+2, r) + diff[l+1] - diff[l]

		// Option 2: use operation with cost x
		cost := solve(l+1, r-1) + x
		if cost < best {
			best = cost
		}

		dp[l][r] = best
		return best
	}

	return solve(0, m-1)
}

func main() {
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("1100011000", "0101001010", 2))
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("10110", "00011", 4))
}
