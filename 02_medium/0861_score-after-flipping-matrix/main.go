package main

// LeetCode #861: Score After Flipping Matrix
// https://leetcode.com/problems/score-after-flipping-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{1, 1}, {1, 1}}))
}

// Time: O(m * n) | Space: O(1)
func ScoreAfterFlippingMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Ensure first column is all 1s
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}

	ans := 0
	for j := 0; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			ones += grid[i][j]
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += ones * (1 << (n - 1 - j))
	}

	return ans
}
