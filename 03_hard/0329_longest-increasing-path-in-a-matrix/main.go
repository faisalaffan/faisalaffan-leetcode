package main

// LeetCode #329: Longest Increasing Path in a Matrix
// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
// Difficulty: Hard

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for r := 0; r < rows; r++ {
		memo[r] = make([]int, cols)
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}
		best := 1
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			path := 1 + dfs(nr, nc)
			if path > best {
				best = path
			}
		}
		memo[r][c] = best
		return best
	}

	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if path := dfs(r, c); path > ans {
				ans = path
			}
		}
	}
	return ans
}

func main() {
	// Example 1
	matrix1 := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Println(longestIncreasingPath(matrix1))
	// 4

	// Example 2
	matrix2 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	fmt.Println(longestIncreasingPath(matrix2))
	// 4

	// Example 3
	matrix3 := [][]int{{1}}
	fmt.Println(longestIncreasingPath(matrix3))
	// 1
}
