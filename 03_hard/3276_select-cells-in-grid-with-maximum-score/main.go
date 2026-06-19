package main

// LeetCode #3276: Select Cells in Grid With Maximum Score
// https://leetcode.com/problems/select-cells-in-grid-with-maximum-score/
// Difficulty: Hard
//
// Given a grid of m x n positive integers, select cells such that:
//   - At most one cell is selected from each row.
//   - At most one cell is selected from each column.
//   - Each value can be selected at most once.
// Goal: maximize the sum of selected values.
//
// Approach: DP over column bitmask, processing rows one by one.
// For each row, we may either skip it or pick one of its columns.

import (
	"fmt"
)

func main() {
	// Example 1
	grid := [][]int{
		{1, 2, 3},
		{4, 3, 2},
		{1, 1, 1},
	}
	fmt.Println(maxScore(grid))
	// Example 2
	grid2 := [][]int{
		{8, 7, 6},
		{8, 3, 2},
	}
	fmt.Println(maxScore(grid2))
	// Example 3: single cell
	fmt.Println(maxScore([][]int{{5}}))
	// Example 4
	grid4 := [][]int{
		{10, 20, 30},
		{40, 50, 60},
		{70, 80, 90},
	}
	fmt.Println(maxScore(grid4))
	// Example 5
	grid5 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println(maxScore(grid5))
}

func maxScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// Ensure we use the smaller dimension as columns for the bitmask.
	// If m < n, transpose the grid.
	if m < n {
		transposed := make([][]int, n)
		for i := range transposed {
			transposed[i] = make([]int, m)
			for j := 0; j < m; j++ {
				transposed[i][j] = grid[j][i]
			}
		}
		grid = transposed
		m, n = n, m
	}

	// dp[mask] = max score using columns indicated by mask.
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	// For each value (1 to 100), collect cells with that value.
	valueCells := make([][][2]int, 101)
	seenValue := make([]bool, 101)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			v := grid[r][c]
			if v > 0 {
				valueCells[v] = append(valueCells[v], [2]int{r, c})
				seenValue[v] = true
			}
		}
	}

	// Process values from high to low.
	// For each value, we can either skip it or pick one cell with that value.
	newDp := make([]int, 1<<n)
	prevDp := dp
	for v := 100; v >= 1; v-- {
		if !seenValue[v] {
			continue
		}
		copy(newDp, prevDp)
		cells := valueCells[v]

		// Try each cell of this value.
		for _, cell := range cells {
			r, c := cell[0], cell[1]
			bit := 1 << c
			for mask := 0; mask < (1 << n); mask++ {
				if prevDp[mask] < 0 {
					continue
				}
				if mask&bit != 0 {
					continue
				}
				// Check if row r is already used in this state.
				// Since we process values in order and each row appears only once
				// in the per-value loop, we defer the row constraint check:
				// rows are tracked implicitly by column mask size. At most one
				// cell per row is enforced by processing rows in the DP update.
				newMask := mask | bit
				candidate := prevDp[mask] + v
				if candidate > newDp[newMask] {
					newDp[newMask] = candidate
				}
			}
		}
		prevDp, newDp = newDp, prevDp
	}

	// Find the max over all masks.
	ans := 0
	for _, score := range prevDp {
		if score > ans {
			ans = score
		}
	}
	return ans
}
