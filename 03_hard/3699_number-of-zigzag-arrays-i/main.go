package main

// LeetCode #3699: Number of ZigZag Arrays I
// https://leetcode.com/problems/number-of-zigzag-arrays-i/
// Difficulty: Hard
//
// Count arrays of length n with values in [l, r] such that:
// 1. No adjacent elements equal
// 2. No three consecutive elements are strictly increasing or decreasing
//
// Approach: DP over possible values with three-state tracking.

import "fmt"

func main() {
	// Example 1
	fmt.Println(zigZagArraysI(3, 4, 5))
	// Example 2
	fmt.Println(zigZagArraysI(4, 1, 3))
	// Edge: n = 3, small range
	fmt.Println(zigZagArraysI(3, 1, 2))
}

const MOD = 1000000007

func zigZagArraysI(n int, l int, r int) int {
	if n < 3 || l > r {
		return 0
	}
	m := r - l + 1
	if m < 2 {
		return 0
	}

	// dp[last][state]
	// state 0: last > second-last (increasing at end)
	// state 1: last < second-last (decreasing at end)
	// state 2: equal doesn't happen (no adjacent equal)
	// Actually since no adjacent equal, we only track up/down

	// For position i, we track counts for each possible value
	dp := make([][2]int, m)
	for v := 0; v < m; v++ {
		dp[v][0] = 1 // increasing (single element)
		dp[v][1] = 1 // decreasing (single element)
	}

	for pos := 2; pos <= n; pos++ {
		ndp := make([][2]int, m)
		for cur := 0; cur < m; cur++ {
			// For arrays where cur is at an odd position (1-indexed from end):
			// array ends with cur, and cur should be a peak or valley
			// We need to consider previous values that are different from cur

			// Previous was less than cur: cur is at a peak
			for prev := 0; prev < cur; prev++ {
				ndp[cur][0] = (ndp[cur][0] + dp[prev][1]) % MOD
			}
			// Previous was greater than cur: cur is at a valley
			for prev := cur + 1; prev < m; prev++ {
				ndp[cur][1] = (ndp[cur][1] + dp[prev][0]) % MOD
			}
		}
		dp = ndp
	}

	result := 0
	for v := 0; v < m; v++ {
		result = (result + dp[v][0] + dp[v][1]) % MOD
	}
	return result
}
