# 1252 — Cells With Odd Values In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func oddCells(m, n int, indices [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m + k)  |  **Ruang:** O(n + m)


## 💻 Solusi Go

```go
package main

// LeetCode #1252: Cells with Odd Values in a Matrix
// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
// Difficulty: Easy
// Time: O(n + m + k) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}})) // 6
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}})) // 0
}

// LeetCode submission: oddCells
func oddCells(m, n int, indices [][]int) int {
  // Alokasi slice
	rows := make([]int, m)
  // Alokasi slice
	cols := make([]int, n)
	for _, idx := range indices {
		rows[idx[0]]++
		cols[idx[1]]++
	}
	oddRows, oddCols := 0, 0
	for _, v := range rows {
		if v%2 == 1 {
			oddRows++
		}
	}
	for _, v := range cols {
		if v%2 == 1 {
			oddCols++
		}
	}
	return oddRows*(n-oddCols) + (m-oddRows)*oddCols
}
```
