package main

// LeetCode #2132: Stamping the Grid
// https://leetcode.com/problems/stamping-the-grid/
// Difficulty: Hard
//
// Approach: 2D prefix sum + difference array.
// 1. Build 2D prefix sum of the grid to query subgrid sums (obstacles = 1).
// 2. For each possible stamp top-left corner, check if the h x w area has sum 0 (no obstacles).
// 3. If so, mark the stamped area using a 2D difference array.
// 4. Reconstruct the coverage from the diff array.
// 5. If any cell with grid[i][j] == 0 has zero coverage, return false.

import "fmt"

func main() {
	// Example from problem statement
	grid1 := [][]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	stampHeight1 := 4
	stampWidth1 := 3
	fmt.Printf("possibleToStamp(...) = %t (expected true)\n",
		possibleToStamp(grid1, stampHeight1, stampWidth1))

	// Additional tests
	grid2 := [][]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Printf("possibleToStamp(...) = %t (expected false)\n",
		possibleToStamp(grid2, 2, 2))

	grid3 := [][]int{{0}}
	fmt.Printf("possibleToStamp(...) = %t (expected true)\n",
		possibleToStamp(grid3, 1, 1))
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	// 2D prefix sum
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + grid[i][j]
		}
	}

	sumRange := func(r1, c1, r2, c2 int) int {
		return prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
	}

	// 2D difference array (size m+1 x n+1 for easy boundary handling)
	diff := make([][]int, m+1)
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	// Mark all valid stamp placements
	for i := 0; i+stampHeight <= m; i++ {
		for j := 0; j+stampWidth <= n; j++ {
			if sumRange(i, j, i+stampHeight-1, j+stampWidth-1) == 0 {
				diff[i][j]++
				diff[i][j+stampWidth]--
				diff[i+stampHeight][j]--
				diff[i+stampHeight][j+stampWidth]++
			}
		}
	}

	// Reconstruct coverage and check
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				diff[i][j] += diff[i-1][j]
			}
			if j > 0 {
				diff[i][j] += diff[i][j-1]
			}
			if i > 0 && j > 0 {
				diff[i][j] -= diff[i-1][j-1]
			}
			if grid[i][j] == 0 && diff[i][j] == 0 {
				return false
			}
		}
	}

	return true
}
