package main

// LeetCode #807: Max Increase to Keep City Skyline
// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > rowMax[i] {
				rowMax[i] = grid[i][j]
			}
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			limit := min(rowMax[i], colMax[j])
			total += limit - grid[i][j]
		}
	}

	return total
}
