package main

// LeetCode #1473: Paint House III
// https://leetcode.com/problems/paint-house-iii/
// Difficulty: Hard
//
// Approach: 3D DP
// dp[i][j][k] = min cost to paint first i houses (0-indexed), where
// house i-1 is painted color j, and there are exactly k neighborhoods.
// Transition:
//   - If houses[i-1] != 0 (already painted), only that color is allowed.
//   - If houses[i-1] == 0, try all colors 1..n.
//   - k increments when current color != previous house's color.
// Answer: min over dp[m][*][target].

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3))
	// Expected: 9

	// Example 2
	fmt.Println(minCost([]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3))
	// Expected: 11

	// Example 3
	fmt.Println(minCost([]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3))
	// Expected: -1 (not possible)
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	const INF = math.MaxInt32

	// dp[i][j][k] — working with 1-indexed for i and k for simplicity
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			for k := range dp[i][j] {
				dp[i][j][k] = INF
			}
		}
	}

	// Base: 0 houses, 0 neighborhoods, any "last color" is 0 cost
	// We'll treat house index 0 (0 houses processed) specially.
	// Actually, let's use 0-based house index and 1-based neighborhood.
	// Simpler: dp[h][c][t] = min cost for first h houses (h from 1..m),
	// last house color c (1..n), exactly t neighborhoods (1..target).
	// Initialize first house.

	prevColor := 0
	for c := 1; c <= n; c++ {
		if houses[0] != 0 && houses[0] != c {
			continue
		}
		paintCost := 0
		if houses[0] == 0 {
			paintCost = cost[0][c-1]
		}
		dp[1][c][1] = paintCost
		prevColor = c
	}
	_ = prevColor

	// Fill DP
	for i := 2; i <= m; i++ {
		for c := 1; c <= n; c++ {
			if houses[i-1] != 0 && houses[i-1] != c {
				continue
			}
			paintCost := 0
			if houses[i-1] == 0 {
				paintCost = cost[i-1][c-1]
			}
			for t := 1; t <= target && t <= i; t++ {
				best := INF
				// Case 1: same color as previous house
				if dp[i-1][c][t] < INF {
					best = minInt(best, dp[i-1][c][t]+paintCost)
				}
				// Case 2: different color from previous house
				if t > 1 {
					for pc := 1; pc <= n; pc++ {
						if pc == c {
							continue
						}
						if dp[i-1][pc][t-1] < INF {
							best = minInt(best, dp[i-1][pc][t-1]+paintCost)
						}
					}
				}
				dp[i][c][t] = best
			}
		}
	}

	ans := INF
	for c := 1; c <= n; c++ {
		if dp[m][c][target] < ans {
			ans = dp[m][c][target]
		}
	}
	if ans == INF {
		return -1
	}
	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
