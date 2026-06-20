package main

// LeetCode #3797: Count Routes to Climb a Rectangular Grid
// https://leetcode.com/problems/count-routes-to-climb-a-rectangular-grid/
// Difficulty: Hard
//
// Count paths from top-left to bottom-right moving only right/down,
// where each step height difference doesn't exceed d.
//
// Approach: DP with prefix sums for efficient transitions.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countRoutes([]string{"10", "20", "30"}, 10))
	// Example 2
	fmt.Println(countRoutes([]string{"5", "5", "5"}, 0))
	// Edge: single cell
	fmt.Println(countRoutes([]string{"7"}, 100))
	// Edge: d=0
	fmt.Println(countRoutes([]string{"1", "2"}, 0))
}

func countRoutes(grid []string, d int) int {
	const mod = 1000000007
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	// Convert grid to ints
	vals := make([][]int, m)
	for i := range vals {
		vals[i] = make([]int, n)
		for j, ch := range grid[i] {
			vals[i][j] = int(ch - '0')
		}
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var ways int
			// From top
			if i > 0 {
				diff := vals[i][j] - vals[i-1][j]
				if diff < 0 {
					diff = -diff
				}
				if diff <= d {
					ways = (ways + dp[i-1][j]) % mod
				}
			}
			// From left
			if j > 0 {
				diff := vals[i][j] - vals[i][j-1]
				if diff < 0 {
					diff = -diff
				}
				if diff <= d {
					ways = (ways + dp[i][j-1]) % mod
				}
			}
			dp[i][j] = ways
		}
	}

	return dp[m-1][n-1]
}
