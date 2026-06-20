# 3651 — Minimum Cost Path With Teleportations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minCost(grid [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3651: Minimum Cost Path with Teleportations
// https://leetcode.com/problems/minimum-cost-path-with-teleportations/
// Difficulty: Hard
//
// Given an m x n grid, go from (0,0) to (m-1,n-1). Normal moves right/down
// cost = destination cell value. You can teleport up to k times from (i,j) to
// any (x,y) where grid[x][y] <= grid[i][j], cost = 0.
//
// Approach: DP with teleport tracking. For each cell, track min cost with
// t teleports used.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(minCost([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 0))
	// Example 2
	fmt.Println(minCost([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 1))
	// Edge: single cell
	fmt.Println(minCost([][]int{{5}}, 0))
}

func minCost(grid [][]int, k int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 0
	}

	// dp[i][j][t] = min cost to reach (i,j) with exactly t teleports used
  // Matriks 2D
	dp := make([][][]int, m)
  // Range loop
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for t := range dp[i][j] {
				dp[i][j][t] = math.MaxInt32
			}
		}
	}

	dp[0][0][0] = 0

	// Process cells in order (top-left to bottom-right)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for t := 0; t <= k; t++ {
				if dp[i][j][t] == math.MaxInt32 {
					continue
				}
				// Move right
				if j+1 < n {
					cost := dp[i][j][t] + grid[i][j+1]
					if cost < dp[i][j+1][t] {
						dp[i][j+1][t] = cost
					}
				}
				// Move down
				if i+1 < m {
					cost := dp[i][j][t] + grid[i+1][j]
					if cost < dp[i+1][j][t] {
						dp[i+1][j][t] = cost
					}
				}
				// Teleport
				if t < k {
					for x := 0; x < m; x++ {
						for y := 0; y < n; y++ {
							if (x == i && y == j) || grid[x][y] > grid[i][j] {
								continue
							}
							if dp[i][j][t] < dp[x][y][t+1] {
								dp[x][y][t+1] = dp[i][j][t]
							}
						}
					}
				}
			}
		}
	}

	result := math.MaxInt32
	for t := 0; t <= k; t++ {
		if dp[m-1][n-1][t] < result {
			result = dp[m-1][n-1][t]
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}
```
