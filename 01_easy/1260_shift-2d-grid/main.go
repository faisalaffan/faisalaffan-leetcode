package main

// LeetCode #1260: Shift 2D Grid
// https://leetcode.com/problems/shift-2d-grid/
// Difficulty: Easy
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	// [[9,1,2],[3,4,5],[6,7,8]]
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	// [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
}

// LeetCode submission: shiftGrid
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			ni, nj := idx/n, idx%n
			ans[ni][nj] = grid[i][j]
		}
	}
	return ans
}
