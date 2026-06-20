# 3446 — Sort Matrix By Diagonals

## Deskripsi

**Soal:** [3446. Sort Matrix By Diagonals](https://leetcode.com/problems/sort-matrix-by-diagonals/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2 log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func sortMatrix(grid [][]int) [][]int`

## Solusi Go

```go
package main

// LeetCode #3446: Sort Matrix by Diagonals
// https://leetcode.com/problems/sort-matrix-by-diagonals/
// Difficulty: Medium
// Time: O(n^2 log n) Space: O(n)

import (
	"fmt"
	"slices"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	if n <= 1 {
		return grid
	}

	// Bottom-left (including main diagonal) — descending
	for i := 0; i < n; i++ {
		r, c := i, 0
		length := n - i
  // Membuat slice untuk menyimpan hasil
		diag := make([]int, length)
		for j := 0; j < length; j++ {
			diag[j] = grid[r+j][c+j]
		}
		slices.SortFunc(diag, func(a, b int) int { return b - a })
		for j := 0; j < length; j++ {
			grid[r+j][c+j] = diag[j]
		}
	}

	// Top-right (excluding main diagonal) — ascending
	for j := 1; j < n; j++ {
		r, c := 0, j
		length := n - j
  // Membuat slice untuk menyimpan hasil
		diag := make([]int, length)
		for k := 0; k < length; k++ {
			diag[k] = grid[r+k][c+k]
		}
		slices.Sort(diag)
		for k := 0; k < length; k++ {
			grid[r+k][c+k] = diag[k]
		}
	}
	return grid
}

func main() {
	fmt.Println(sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}))
	// [[8 2 3] [9 6 7] [4 5 1]]
	fmt.Println(sortMatrix([][]int{{0, 1}, {2, 3}}))
	// [[2 1] [2 3]]
}
```
