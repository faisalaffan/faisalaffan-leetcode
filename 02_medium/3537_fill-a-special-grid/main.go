package main

// LeetCode #3537: Fill a Special Grid
// https://leetcode.com/problems/fill-a-special-grid/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	FillASpecialGrid(grid)
	fmt.Println("Test 1:", grid)
	// Test case 2
	grid2 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	FillASpecialGrid(grid2)
	fmt.Println("Test 2:", grid2)
	// Test case 3
	grid3 := [][]int{{0}}
	FillASpecialGrid(grid3)
	fmt.Println("Test 3:", grid3)
}

func FillASpecialGrid(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	m, n := len(grid), len(grid[0])
	// Fill each cell with the sum of its row and column index
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = i + j
			}
		}
	}
}
