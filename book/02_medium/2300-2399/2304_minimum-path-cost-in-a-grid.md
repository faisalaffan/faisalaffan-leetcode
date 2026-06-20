# 2304 — Minimum Path Cost In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minPathCost(grid [][]int, moveCost [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m * n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2304: Minimum Path Cost in a Grid
// https://leetcode.com/problems/minimum-path-cost-in-a-grid/
// Difficulty: Medium
// Time: O(m * n^2) | Space: O(n)

import "fmt"

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice
	dp := make([]int, n)
	copy(dp, grid[0])

	for r := 1; r < m; r++ {
  // Alokasi slice
		next := make([]int, n)
		for j := 0; j < n; j++ {
			next[j] = 1 << 30
			for k := 0; k < n; k++ {
				cost := dp[k] + moveCost[grid[r-1][k]][j] + grid[r][j]
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	minCost := dp[0]
	for _, v := range dp {
		if v < minCost {
			minCost = v
		}
	}
	return minCost
}

func main() {
	// Test case 1
	fmt.Println(minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}))
	// Expected: 17

	// Test case 2
	fmt.Println(minPathCost([][]int{{5, 1, 2}, {4, 0, 3}}, [][]int{{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2}}))
	// Expected: 6
}
```
