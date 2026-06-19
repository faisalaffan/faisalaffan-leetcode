package main

// LeetCode #688: Knight Probability in Chessboard
// https://leetcode.com/problems/knight-probability-in-chessboard/
// Difficulty: Medium
// Time: O(K * N^2)
// Space: O(N^2)

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(1, 0, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, n)
	}
	dp[row][column] = 1.0

	for step := 0; step < k; step++ {
		next := make([][]float64, n)
		for i := range next {
			next[i] = make([]float64, n)
		}
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if dp[r][c] == 0 {
					continue
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < n {
						next[nr][nc] += dp[r][c] / 8.0
					}
				}
			}
		}
		dp = next
	}

	result := 0.0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			result += dp[r][c]
		}
	}
	return result
}
