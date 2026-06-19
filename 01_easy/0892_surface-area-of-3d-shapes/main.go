package main

// LeetCode #892: Surface Area of 3D Shapes
// https://leetcode.com/problems/surface-area-of-3d-shapes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{{2}}))                                   // 10
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))                       // 34
}

// surfaceArea calculates the surface area of a 3D shape on a grid.
// Time: O(n^2). Space: O(1).
func surfaceArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				// Top + bottom
				area += 2
				// Four sides - subtract adjacent overlaps
				area += grid[i][j] * 4
				if i > 0 {
					area -= min(grid[i][j], grid[i-1][j]) * 2
				}
				if j > 0 {
					area -= min(grid[i][j], grid[i][j-1]) * 2
				}
			}
		}
	}
	return area
}
