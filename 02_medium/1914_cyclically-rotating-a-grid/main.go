package main

// LeetCode #1914: Cyclically Rotating a Grid
// https://leetcode.com/problems/cyclically-rotating-a-grid/
// Difficulty: Medium

import "fmt"

func main() {
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(RotateGrid(grid, 2))

	grid2 := [][]int{{10, 20}, {30, 40}}
	fmt.Println(RotateGrid(grid2, 1))
}

// Time: O(m*n), Space: O(m+n)
func RotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top < bottom && left < right {
		layerLen := 2*(right-left+1) + 2*(bottom-top+1) - 4
		kMod := k % layerLen

		// Extract layer
		arr := make([]int, layerLen)
		idx := 0
		for j := left; j <= right; j++ {
			arr[idx] = grid[top][j]
			idx++
		}
		for i := top + 1; i <= bottom; i++ {
			arr[idx] = grid[i][right]
			idx++
		}
		for j := right - 1; j >= left; j-- {
			arr[idx] = grid[bottom][j]
			idx++
		}
		for i := bottom - 1; i > top; i-- {
			arr[idx] = grid[i][left]
			idx++
		}

		// Place rotated
		idx = 0
		for j := left; j <= right; j++ {
			grid[top][j] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for i := top + 1; i <= bottom; i++ {
			grid[i][right] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for j := right - 1; j >= left; j-- {
			grid[bottom][j] = arr[(idx+kMod)%layerLen]
			idx++
		}
		for i := bottom - 1; i > top; i-- {
			grid[i][left] = arr[(idx+kMod)%layerLen]
			idx++
		}

		top++
		bottom--
		left++
		right--
	}
	return grid
}
