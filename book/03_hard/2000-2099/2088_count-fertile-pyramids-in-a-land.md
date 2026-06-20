# 2088 — Count Fertile Pyramids In A Land

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countPyramids(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2088: Count Fertile Pyramids in a Land
// https://leetcode.com/problems/count-fertile-pyramids-in-a-land/
// Difficulty: Hard
//
// DP approach: dp[i][j] = max pyramid height with (i,j) as the top.
// Regular pyramid (top down): dp[i][j] = 1 + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
// Inverted pyramid (top up):  dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
// Sum (dp[i][j] - 1) over all cells where dp[i][j] > 1.

import "fmt"

func main() {
	fmt.Println(countPyramids([][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1}}))
	fmt.Println(countPyramids([][]int{{1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

func countPyramids(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Matriks 2D
	dp := make([][]int, m)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0

	// Regular pyramids (top pointing down): bottom-up DP
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == m-1 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i+1][j-1], min(dp[i+1][j], dp[i+1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	// Inverted pyramids (top pointing up): top-down DP
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i-1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	return total
}
```
