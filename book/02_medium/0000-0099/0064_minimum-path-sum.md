# 0064 — Minimum Path Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minPathSum(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #64: Minimum Path Sum
// https://leetcode.com/problems/minimum-path-sum/
// Difficulty: Medium

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 7

	// Test case 2
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})) // 12

	// Test case 3
	fmt.Println(minPathSum([][]int{{1}})) // 1
}

// Time: O(m*n) | Space: O(1)
```
