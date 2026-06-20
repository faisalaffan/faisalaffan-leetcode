# 1895 — Largest Magic Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func LargestMagicSquare(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(m*n*min(m,n)^2), Space: O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1895: Largest Magic Square
// https://leetcode.com/problems/largest-magic-square/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LargestMagicSquare([][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}))
	fmt.Println(LargestMagicSquare([][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}}))
}

// Time: O(m*n*min(m,n)^2), Space: O(m*n)
func LargestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Build prefix sums for rows and columns
  // Matriks 2D
	rowSum := make([][]int, m+1)
  // Matriks 2D
	colSum := make([][]int, m+1)
  // Range loop
	for i := range rowSum {
		rowSum[i] = make([]int, n+1)
		colSum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSum[i+1][j+1] = rowSum[i+1][j] + grid[i][j]
			colSum[i+1][j+1] = colSum[i][j+1] + grid[i][j]
		}
	}

	maxK := min(m, n)
	for k := maxK; k >= 1; k-- {
		for i := 0; i+k <= m; i++ {
			for j := 0; j+k <= n; j++ {
				if isMagic(grid, i, j, k, rowSum, colSum) {
					return k
				}
			}
		}
	}
	return 1
}

func isMagic(grid [][]int, r, c, k int, rowSum, colSum [][]int) bool {
	target := rowSum[r+1][c+k] - rowSum[r+1][c]

	// Check rows
	for i := 0; i < k; i++ {
		sum := rowSum[r+i+1][c+k] - rowSum[r+i+1][c]
		if sum != target {
			return false
		}
	}
	// Check columns
	for j := 0; j < k; j++ {
		sum := colSum[r+k][c+j+1] - colSum[r][c+j+1]
		if sum != target {
			return false
		}
	}
	// Check diagonals
	diag1, diag2 := 0, 0
	for i := 0; i < k; i++ {
		diag1 += grid[r+i][c+i]
		diag2 += grid[r+i][c+k-1-i]
	}
	return diag1 == target && diag2 == target
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
