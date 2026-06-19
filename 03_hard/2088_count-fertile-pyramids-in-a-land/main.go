package main

// LeetCode #2088: Count Fertile Pyramids in a Land
// https://leetcode.com/problems/count-fertile-pyramids-in-a-land/
// Difficulty: Hard
//
// Approach: DP. For each fertile cell, dp[i][j] = max pyramid height with (i,j) as top/inverted top.
// Recurrence (regular pyramid): dp[i][j] = 1 + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
// Recurrence (inverted pyramid): dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
// Total pyramids = sum(dp[i][j] - 1) over all cells where dp[i][j] > 1.

import "fmt"

func main() {
	// Example from problem statement
	grid1 := [][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}
	fmt.Printf("countPyramids(%v) = %d (expected 2)\n", grid1, countPyramids(grid1))

	// Additional test cases
	grid2 := [][]int{{1, 1, 1}, {1, 1, 1}}
	fmt.Printf("countPyramids(%v) = %d\n", grid2, countPyramids(grid2))

	grid3 := [][]int{{1}}
	fmt.Printf("countPyramids(%v) = %d (expected 0)\n", grid3, countPyramids(grid3))

	grid4 := [][]int{{1, 1}}
	fmt.Printf("countPyramids(%v) = %d (expected 0)\n", grid4, countPyramids(grid4))

	grid5 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Printf("countPyramids(%v) = %d\n", grid5, countPyramids(grid5))
}

func countPyramids(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Regular pyramids (top pointing down): bottom-up DP
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == m-1 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i+1][j-1], min(dp[i+1][j], dp[i+1][j+1]))
			}
		}
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	// Inverted pyramids (top pointing up): top-down DP
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i-1][j+1]))
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	return total
}
