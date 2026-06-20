# 0063 — Unique Paths Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func uniquePathsWithObstacles(obstacleGrid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #63: Unique Paths II
// https://leetcode.com/problems/unique-paths-ii/
// Difficulty: Medium

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

  // Matriks 2D
	dp := make([][]int, m)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})) // 2

	// Test case 2
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}})) // 1

	// Test case 3
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}})) // 0
}

// Time: O(m*n) | Space: O(m*n)
```
