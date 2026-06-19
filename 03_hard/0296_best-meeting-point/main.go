package main

// LeetCode #296: Best Meeting Point
// https://leetcode.com/problems/best-meeting-point/
// Difficulty: Hard [Paid]
//
// Approach: Median of coordinates (Manhattan distance).
//   The total Manhattan distance is minimized at the median of row coordinates
//   and median of column coordinates independently.
//   1. Collect all row coordinates and column coordinates where grid[i][j] == 1.
//   2. Sort each list.
//   3. The optimal meeting point is at the median values.
//   4. Sum the absolute differences from the medians.

import (
	"fmt"
	"sort"
)

func main() {
	// Example: [[1,0,0,0,1],[0,0,0,0,0],[0,0,1,0,0]]
	grid := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println("Best meeting point distance:", minTotalDistance(grid))

	// Test: [[0,0],[2,0],[1,1]]
	grid2 := [][]int{
		{0, 0},
		{2, 0},
		{1, 1},
	}
	fmt.Println("Best meeting point distance (2):", minTotalDistance(grid2)) // 4

	// Test: [[1,1]]
	grid3 := [][]int{
		{1, 1},
	}
	fmt.Println("Single row with 2 points:", minTotalDistance(grid3))

	// Test: no friends (should return 0).
	grid4 := [][]int{
		{0, 0},
		{0, 0},
	}
	fmt.Println("No friends:", minTotalDistance(grid4))
}

// minTotalDistance returns the minimum total Manhattan distance for a meeting point.
func minTotalDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Collect row and column coordinates of all 1s.
	var rows, cols []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	if len(rows) == 0 {
		return 0
	}

	// Sort both lists.
	sort.Ints(rows)
	sort.Ints(cols)

	// Find medians.
	medianRow := rows[len(rows)/2]
	medianCol := cols[len(cols)/2]

	// Sum distances from median.
	total := 0
	for _, r := range rows {
		total += abs(r - medianRow)
	}
	for _, c := range cols {
		total += abs(c - medianCol)
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Stub compatibility.
func BestMeetingPoint() any {
	grid := [][]int{
		{0, 0},
		{2, 0},
		{1, 1},
	}
	return minTotalDistance(grid)
}
