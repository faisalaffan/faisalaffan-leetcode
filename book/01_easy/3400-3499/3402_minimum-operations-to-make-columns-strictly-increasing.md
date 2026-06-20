# 3402 — Minimum Operations To Make Columns Strictly Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MinimumOperationsToMakeColumnsStrictlyIncreasing(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3402: Minimum Operations to Make Columns Strictly Increasing
// https://leetcode.com/problems/minimum-operations-to-make-columns-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
	fmt.Println(MinimumOperationsToMakeColumnsStrictlyIncreasing([][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}}))
}

// MinimumOperationsToMakeColumnsStrictlyIncreasing returns minimum operations to make each column strictly increasing.
// Each operation increments an element by 1.
// Time: O(n * m). Space: O(1).
func MinimumOperationsToMakeColumnsStrictlyIncreasing(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	ops := 0
	for c := 0; c < cols; c++ {
		for r := 1; r < rows; r++ {
			if grid[r][c] <= grid[r-1][c] {
				diff := grid[r-1][c] - grid[r][c] + 1
				ops += diff
				grid[r][c] += diff
			}
		}
	}
	return ops
}
```
