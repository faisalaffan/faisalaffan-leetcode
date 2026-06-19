package main

// LeetCode #576: Out of Boundary Paths
// https://leetcode.com/problems/out-of-boundary-paths/
// Difficulty: Medium
// Time: O(m * n * maxMove)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(FindPaths(2, 2, 2, 0, 0))
	fmt.Println(FindPaths(1, 3, 3, 0, 1))
}

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	const mod = 1_000_000_007
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	total := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for move := 1; move <= maxMove; move++ {
		next := make([][]int, m)
		for i := range next {
			next[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] == 0 {
					continue
				}
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n {
						total = (total + dp[i][j]) % mod
					} else {
						next[ni][nj] = (next[ni][nj] + dp[i][j]) % mod
					}
				}
			}
		}
		dp = next
	}

	return total
}
