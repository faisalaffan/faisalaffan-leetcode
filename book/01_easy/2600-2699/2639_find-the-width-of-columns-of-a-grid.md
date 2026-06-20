# 2639 — Find The Width Of Columns Of A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2639: Find the Width of Columns of a Grid
// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/
// Difficulty: Easy
// Time: O(m * n) | Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{1}, {22}, {333}}))
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}))
}

func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		maxLen := 0
		for i := 0; i < m; i++ {
			width := len(strconv.Itoa(grid[i][j]))
			if width > maxLen {
				maxLen = width
			}
		}
		ans[j] = maxLen
	}
	return ans
}
```
